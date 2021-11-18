package service

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"

	"github.com/free5gc/http2_util"
	"github.com/free5gc/logger_util"
	"github.com/free5gc/path_util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"samf.com/consumer"
	samf_context "samf.com/context"
	"samf.com/factory"
	"samf.com/logger"
	"samf.com/samf_service"
	"samf.com/util"
)

type SAMF struct{}

type (
	// config information
	Config struct {
		samfcfg string
	}
)

var config Config

var samfCLi = []cli.Flag{
	cli.StringFlag{
		Name:  "free5gccfg",
		Usage: "common config file",
	},
	cli.StringFlag{
		Name:  "samfcfg",
		Usage: "samf config file",
	},
}

var initLog *logrus.Entry

func init() {
	initLog = logger.InitLog
}

func (*SAMF) GetCliCmd() (flags []cli.Flag) {
	return samfCLi
}

func (samf *SAMF) Initialize(c *cli.Context) error {
	config = Config{
		samfcfg: c.String("samfcfg"),
	}

	if config.samfcfg != "" {
		if err := factory.InitConfigFactory(config.samfcfg); err != nil {
			return err
		}
	} else {
		DefaultSAMFConfigPath := path_util.Free5gcPath("free5gc/config/samfcfg.yaml")
		if err := factory.InitConfigFactory(DefaultSAMFConfigPath); err != nil {
			return err
		}
	}

	samf.setLogLevel()

	if err := factory.CheckConfigVersion(); err != nil {
		return err
	}

	return nil
}

func (samf *SAMF) setLogLevel() {
	if factory.SamfConfig.Logger == nil {
		initLog.Warnln("SAMF config without log level setting!!!")
		return
	}

	logger.SetLogLevel(logrus.InfoLevel)

}

func (samf *SAMF) FilterCli(c *cli.Context) (args []string) {
	for _, flag := range samf.GetCliCmd() {
		name := flag.GetName()
		value := fmt.Sprint(c.Generic(name))
		if value == "" {
			continue
		}

		args = append(args, "--"+name, value)
	}
	return args
}

func (samf *SAMF) Start() {

	initLog.Infoln("Server started")

	if !util.InitSAMFContext() {
		initLog.Error("Initicating context failed")
		return
	}

	wg := sync.WaitGroup{}

	self := samf_context.SAMF_Self()
	util.InitSamfContext(self)

	addr := fmt.Sprintf("127.0.0.1:24243")
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	samf_service.AddService(router)

	profile := consumer.BuildNFInstance(self)
	var newNrfUri string
	var err error

	newNrfUri, self.NfId, err = consumer.SendRegisterNFInstance(self.NrfUri, profile.NfInstanceId, profile)
	if err == nil {
		self.NrfUri = newNrfUri
	} else {
		initLog.Errorf("Send Register NFInstance Error[%s]", err.Error())
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		os.Exit(0)
	}()

	server, err := http2_util.NewServer(addr, "samfsslkey.log", router)
	if server == nil {
		initLog.Errorf("Initialize HTTP server failed: %+v", err)
		return
	}
	if err != nil {
		initLog.Warnln("Initialize HTTP server:", err)
	}
	serverScheme := factory.SamfConfig.Configuration.Sbi.Scheme
	if serverScheme == "http" {
		err = server.ListenAndServe()
	} else if serverScheme == "https" {
		err = server.ListenAndServe() //TODO: changing to HTTPS (TLS)
	}

	if err != nil {
		initLog.Fatalln("HTTP server setup failed:", err)
	}
	initLog.Info("SAMF running...")

	wg.Wait()
}

func (samf *SAMF) Exec(c *cli.Context) error {

	initLog.Traceln("args:", c.String("samfcfg"))
	args := samf.FilterCli(c)
	initLog.Traceln("filter: ", args)
	command := exec.Command("./samf", args...)

	wg := sync.WaitGroup{}
	wg.Add(3)

	stdout, err := command.StdoutPipe()
	if err != nil {
		initLog.Fatalln(err)
	}
	go func() {
		in := bufio.NewScanner(stdout)
		for in.Scan() {
			fmt.Println(in.Text())
		}
		wg.Done()
	}()

	stderr, err := command.StderrPipe()
	if err != nil {
		initLog.Fatalln(err)
	}
	go func() {
		in := bufio.NewScanner(stderr)
		for in.Scan() {
			fmt.Println(in.Text())
		}
		wg.Done()
	}()

	go func() {
		if errCom := command.Start(); errCom != nil {
			initLog.Errorf("SAMF start error: %v", errCom)
		}
		wg.Done()
	}()

	wg.Wait()

	return err
}
