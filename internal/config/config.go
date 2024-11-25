package config

import (
	"os"
)

type Config struct {
	Port   string
	DBURL  string
	APIURL string
}

func GetConfig() *Config {
	port, _ := os.LookupEnv("PORT")
	dburl, _ := os.LookupEnv("DB_URL")
	apiurl, _ := os.LookupEnv("GET_INFO")

	return &Config{
		Port:   port,
		DBURL:  dburl,
		APIURL: apiurl,
	}
}
