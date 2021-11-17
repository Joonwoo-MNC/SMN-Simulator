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
	"ssmf.com/consumer"
	ssmf_context "ssmf.com/context"
	"ssmf.com/factory"
	"ssmf.com/logger"
	"ssmf.com/util"
	"ssmf.com/ssmf_service"
)

type SSMF struct{}

type (
	// config information
	Config struct {
		ssmfcfg string
	}
)

var config Config

var ssmfCLi = []cli.Flag{
	cli.StringFlag{
		Name:  "free5gccfg",
		Usage: "common config file",
	},
	cli.StringFlag{
		Name:  "ssmfcfg",
		Usage: "ssmf config file",
	},
}

var initLog *logrus.Entry

func init() {
	initLog = logger.InitLog
}

func (*SSMF) GetCliCmd() (flags []cli.Flag) {
	return ssmfCLi
}

func (ssmf *SSMF) Initialize(c *cli.Context) error {
	config = Config{
		ssmfcfg: c.String("ssmfcfg"),
	}

	if config.ssmfcfg != "" {
		if err := factory.InitConfigFactory(config.ssmfcfg); err != nil {
			return err
		}
	} else {
		DefaultSSMFConfigPath := path_util.Free5gcPath("free5gc/config/ssmfcfg.yaml")
		if err := factory.InitConfigFactory(DefaultSSMFConfigPath); err != nil {
			return err
		}
	}

	ssmf.setLogLevel()

	if err := factory.CheckConfigVersion(); err != nil {
		return err
	}

	return nil
}

func (ssmf *SSMF) setLogLevel() {
	if factory.SsmfConfig.Logger == nil {
		initLog.Warnln("SSMF config without log level setting!!!")
		return
	}

	logger.SetLogLevel(logrus.InfoLevel)

}

func (ssmf *SSMF) FilterCli(c *cli.Context) (args []string) {
	for _, flag := range ssmf.GetCliCmd() {
		name := flag.GetName()
		value := fmt.Sprint(c.Generic(name))
		if value == "" {
			continue
		}

		args = append(args, "--"+name, value)
	}
	return args
}

func (ssmf *SSMF) Start() {

	initLog.Infoln("Server started")

	if !util.InitSSMFContext() {
		initLog.Error("Initicating context failed")
		return
	}

	wg := sync.WaitGroup{}

	self := ssmf_context.SSMF_Self()
	util.InitSsmfContext(self)

	addr := fmt.Sprintf("127.0.0.1:24244")
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	ssmf_service.AddService(router)

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

	server, err := http2_util.NewServer(addr, "ssmfsslkey.log", router)
	if server == nil {
		initLog.Errorf("Initialize HTTP server failed: %+v", err)
		return
	}
	if err != nil {
		initLog.Warnln("Initialize HTTP server:", err)
	}
	serverScheme := factory.SsmfConfig.Configuration.Sbi.Scheme
	if serverScheme == "http" {
		err = server.ListenAndServe()
	} else if serverScheme == "https" {
		err = server.ListenAndServe() //TODO: changing to HTTPS (TLS)
	}

	if err != nil {
		initLog.Fatalln("HTTP server setup failed:", err)
	}
	initLog.Info("SSMF running...")

	wg.Wait()
}

func (ssmf *SSMF) Exec(c *cli.Context) error {

	initLog.Traceln("args:", c.String("ssmfcfg"))
	args := ssmf.FilterCli(c)
	initLog.Traceln("filter: ", args)
	command := exec.Command("./ssmf", args...)

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
			initLog.Errorf("SSMF start error: %v", errCom)
		}
		wg.Done()
	}()

	wg.Wait()

	return err
}
