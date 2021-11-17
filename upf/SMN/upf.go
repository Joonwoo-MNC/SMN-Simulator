package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/free5gc/version"
	"smn.com/logger"
	"smn.com/service"
)

// func main() {
// 	message := service.Hello("servicehello")
// 	fmt.Println(message)
// 	message2 := logger.Hello("loggerhello")
// 	fmt.Println(message2)
// 	message3 := factory.Hello("factoryhello")
// 	fmt.Println(message3)

// }

var SMN = &service.SMN{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "smn"
	appLog.Infoln(app.Name)
	appLog.Infoln("SMN version: ", version.GetVersion())
	app.Usage = "-free5gccfg common configuration file -smncfg smn configuration file"
	app.Action = action
	app.Flags = SMN.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		appLog.Errorf("SMN Run error: %v", err)
		return
	}
}

func action(c *cli.Context) error {
	if err := SMN.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("Failed to initialize !!")
	}

	SMN.Start()

	return nil
}
