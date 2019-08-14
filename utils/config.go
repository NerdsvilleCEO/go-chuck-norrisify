package utils

import (
	"os"

	"github.com/go-ini/ini"
)

type AppConfig struct {
	File *ini.File
	Port string
}

func LoadConfig() *AppConfig {
	cfg, err := ini.Load("config.ini")
	PanicIf(err)
	appConfig := &AppConfig{
		File: cfg,
	}
	appConfig.Port = appConfig.getPort()
	return appConfig
}

func getSettings(cf *AppConfig, section string) *ini.Section {
	s, _ := cf.File.GetSection(section)
	return s
}

func (cf *AppConfig) getPort() string {
	s := getSettings(cf, "app")

	if s != nil && s.HasKey("port") {
		port, _ := s.GetKey("port")
		return port.String()
	}

	// Convention for gin, if no port defined in config, use PORT
	port := os.Getenv("PORT")
	if len(port) == 0 {
		panic("You must provide a port to serve this API")
	}

	return port
}
