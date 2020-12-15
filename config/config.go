package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DB     DbConfiguration
	Server ServerConfiguration
}

type Env string

func (e Env) IsLocal() bool {
	return e == "local"
}

func (e Env) IsDev() bool {
	return e == "development"
}

func (e Env) IsDocker() bool {
	return e == "docker"
}

func (e Env) IsProd() bool {
	return e == "production"
}

var config Configuration

func GetConfig() Configuration {
	env := os.Getenv("SOROHA_ENV")
	if "" == env {
		env = "development"
	}
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	dbConfig := DbConfiguration{
		DbURL: os.Getenv("DB_URL"),
	}
	serverConfig := ServerConfiguration{
		PORT: os.Getenv("SERVER_PORT"),
		KEY:  os.Getenv("SECRET_KEY"),
	}

	config = Configuration{
		DB:     dbConfig,
		Server: serverConfig,
	}
	return config
}
