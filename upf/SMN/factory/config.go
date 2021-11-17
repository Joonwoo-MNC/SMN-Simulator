/*
 * SMN Configuration Factory
 */

package factory

import (
	"github.com/free5gc/logger_util"
)

const (
	SMN_EXPECTED_CONFIG_VERSION = "1.0.1"
)

type Config struct {
	Info          *Info               `yaml:"info"`
	Configuration *Configuration      `yaml:"configuration"`
	Logger        *logger_util.Logger `yaml:"logger"`
	Sbi           *Sbi                `yaml:"sbi,omitempty"`
	NrfUri        string              `yaml:"nrfUri,omitempty"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
}

const (
	SMN_DEFAULT_IPV4     = "127.0.0.38"
	SMN_DEFAULT_PORT     = "8000"
	SMN_DEFAULT_PORT_INT = 8000
	SMN_DEFAULT_NRFURI   = "https://127.0.0.10:8000"
)

type Configuration struct {
	SmnName string `yaml:"smnName,omitempty"`
	Sbi       *Sbi   `yaml:"sbi,omitempty"`
	NrfUri    string `yaml:"nrfUri,omitempty"`
}

func (c *Config) GetVersion() string {
	if c.Info != nil && c.Info.Version != "" {
		return c.Info.Version
	}
	return ""
}

type Sbi struct {
	Scheme       string `yaml:"scheme"`
	RegisterIPv4 string `yaml:"registerIPv4,omitempty"` // IP that is registered at NRF.
	BindingIPv4  string `yaml:"bindingIPv4,omitempty"`  // IP used to run the server in the node.
	Port         int    `yaml:"port,omitempty"`
}
