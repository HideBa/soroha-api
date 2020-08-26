package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type Configuration struct {
	DB     DbConfiguration
	Server ServerConfiguration
}

var config Configuration

func GetConfig() Configuration {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("failure to load file: %v", err)
		os.Exit(1)
	}
	dbConfig := DbConfiguration{
		DbName: cfg.Section("database").Key("db_name").String(),
		DbUser: cfg.Section("database").Key("db_user").String(),
		DbPass: cfg.Section("database").Key("db_pass").String(),
		DbHost: cfg.Section("database").Key("db_host").String(),
		DbPort: cfg.Section("database").Key("db_port").String(),
	}
	serverConfig := ServerConfiguration{
		PORT: cfg.Section("server").Key("port").String(),
		KEY:  cfg.Section("server").Key("secret").String(),
	}

	config = Configuration{
		DB:     dbConfig,
		Server: serverConfig,
	}
	return config
}
