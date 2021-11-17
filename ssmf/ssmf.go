package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/free5gc/version"
	"ssmf.com/logger"
	"ssmf.com/service"
)

// func main() {
// 	message := service.Hello("servicehello")
// 	fmt.Println(message)
// 	message2 := logger.Hello("loggerhello")
// 	fmt.Println(message2)
// 	message3 := factory.Hello("factoryhello")
// 	fmt.Println(message3)

// }

var SSMF = &service.SSMF{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "ssmf"
	appLog.Infoln(app.Name)
	appLog.Infoln("SSMF version: ", version.GetVersion())
	app.Usage = "-free5gccfg common configuration file -ssmfcfg ssmf configuration file"
	app.Action = action
	app.Flags = SSMF.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		appLog.Errorf("SSMF Run error: %v", err)
		return
	}
}

func action(c *cli.Context) error {
	if err := SSMF.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("Failed to initialize !!")
	}

	SSMF.Start()

	return nil
}
