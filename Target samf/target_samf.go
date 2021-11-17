package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/free5gc/version"
	"samf.com/logger"
	"samf.com/service"
)

// func main() {
// 	message := service.Hello("servicehello")
// 	fmt.Println(message)
// 	message2 := logger.Hello("loggerhello")
// 	fmt.Println(message2)
// 	message3 := factory.Hello("factoryhello")
// 	fmt.Println(message3)

// }

var SAMF = &service.SAMF{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "samf"
	appLog.Infoln(app.Name)
	appLog.Infoln("SAMF version: ", version.GetVersion())
	app.Usage = "-free5gccfg common configuration file -samfcfg samf configuration file"
	app.Action = action
	app.Flags = SAMF.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		appLog.Errorf("SAMF Run error: %v", err)
		return
	}
}

func action(c *cli.Context) error {
	if err := SAMF.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("Failed to initialize !!")
	}

	SAMF.Start()

	return nil
}
