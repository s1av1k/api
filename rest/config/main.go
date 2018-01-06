package config

import (
	"fmt"
)

type Configuration struct {
	ConfigFile string `yaml:"ConfigFile"`
	URISchema  string
	CertFile   string `yaml:"CertFile"`
	KeyFile    string `yaml:"KeyFile"`
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	GlobalPort int
	Verbose    bool   `yaml:"Verbose"`
}

var Server  Configuration
var Default Configuration

func InitDefault(c *Configuration) {
	c.ConfigFile = ""
	c.URISchema  = "https://"
	c.CertFile   = "server.crt"
	c.KeyFile    = "server.key"
	c.Host       = "localhost"
	c.Port       = 64443
	c.GlobalPort = 443
	c.Verbose    = false
}

func ServerURI(uriPath string) string {
	var scheme string

	if Default.GlobalPort == Server.Port {
		scheme = Server.Host
	} else {
		scheme = fmt.Sprintf("%s:%d", Server.Host, Server.Port)
	}

	return Server.URISchema + scheme + uriPath
}

func Init() {
	InitDefault(&Default)

	InitCommand(&Server)
	if Server == Default {
		InitEnvironment(&Server)
	}
	InitFile(&Server)

	Server.URISchema = Default.URISchema
	Server.GlobalPort = Default.GlobalPort
}