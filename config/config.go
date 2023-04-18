package config

import (
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	ConnectionString string
	Port             string
}

var config appConfig

func Init() (appConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return appConfig{}, err
	}
	config.ConnectionString = os.Getenv("DB_CONNECTION")
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	return config, nil
}
