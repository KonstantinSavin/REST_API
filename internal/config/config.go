package config

import (
	"os"
)

type Config struct {
	Port     string
	LogLevel string
	DBURL    string
}

func GetConfig() *Config {
	port, _ := os.LookupEnv("PORT")
	logl, _ := os.LookupEnv("LOG_LEVEL")
	dburl, _ := os.LookupEnv("DB_URL")

	return &Config{
		Port:     port,
		LogLevel: logl,
		DBURL:    dburl,
	}
}
