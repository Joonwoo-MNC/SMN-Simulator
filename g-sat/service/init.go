package service

import (
	"bufio"
	"fmt"
	"os/exec"
	"sync"

	"github.com/free5gc/http2_util"
	"github.com/free5gc/logger_util"
	"github.com/free5gc/path_util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"smn.com/factory"
	"smn.com/logger"
	"smn.com/smn_service"
)

type SMN struct{}

type (
	// config information
	Config struct {
		smncfg string
	}
)

var config Config

var smnCLi = []cli.Flag{
	cli.StringFlag{
		Name:  "free5gccfg",
		Usage: "common config file",
	},
	cli.StringFlag{
		Name:  "smncfg",
		Usage: "smn config file",
	},
}

var initLog *logrus.Entry

func init() {
	initLog = logger.InitLog
}

func (*SMN) GetCliCmd() (flags []cli.Flag) {
	return smnCLi
}

func (smn *SMN) Initialize(c *cli.Context) error {
	config = Config{
		smncfg: c.String("smncfg"),
	}

	if config.smncfg != "" {
		if err := factory.InitConfigFactory(config.smncfg); err != nil {
			return err
		}
	} else {
		DefaultSMNConfigPath := path_util.Free5gcPath("free5gc/config/smncfg.yaml")
		if err := factory.InitConfigFactory(DefaultSMNConfigPath); err != nil {
			return err
		}
	}

	smn.setLogLevel()

	if err := factory.CheckConfigVersion(); err != nil {
		return err
	}

	return nil
}

func (smn *SMN) setLogLevel() {
	if factory.SmnConfig.Logger == nil {
		initLog.Warnln("SMN config without log level setting!!!")
		return
	}

	logger.SetLogLevel(logrus.InfoLevel)

}

func (smn *SMN) FilterCli(c *cli.Context) (args []string) {
	for _, flag := range smn.GetCliCmd() {
		name := flag.GetName()
		value := fmt.Sprint(c.Generic(name))
		if value == "" {
			continue
		}

		args = append(args, "--"+name, value)
	}
	return args
}

func (smn *SMN) Start() {

	initLog.Infoln("Server started")
	wg := sync.WaitGroup{}

	addr := fmt.Sprintf("127.0.0.1:24247")
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	smn_service.AddService(router)

	server, err := http2_util.NewServer(addr, "smnsslkey.log", router)
	if server == nil {
		initLog.Errorf("Initialize HTTP server failed: %+v", err)
		return
	}
	if err != nil {
		initLog.Warnln("Initialize HTTP server:", err)
	}

	serverScheme := factory.SmnConfig.Configuration.Sbi.Scheme
	if serverScheme == "http" {
		err = server.ListenAndServe()
	} else if serverScheme == "https" {
		err = server.ListenAndServe() //TODO: changing to HTTPS (TLS)
	}

	if err != nil {
		initLog.Fatalln("HTTP server setup failed:", err)
	}
	initLog.Info("SMN running...")

	wg.Wait()

}

func (smn *SMN) Exec(c *cli.Context) error {

	initLog.Traceln("args:", c.String("smncfg"))
	args := smn.FilterCli(c)
	initLog.Traceln("filter: ", args)
	command := exec.Command("./smn", args...)

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
			initLog.Errorf("SMN start error: %v", errCom)
		}
		wg.Done()
	}()

	wg.Wait()

	return err
}
