package config

import (
	"os"
)

type Config struct {
	Port  string
	DBURL string
}

func GetConfig() *Config {
	port, _ := os.LookupEnv("PORT")
	dburl, _ := os.LookupEnv("DB_URL")

	return &Config{
		Port:  port,
		DBURL: dburl,
	}
}
