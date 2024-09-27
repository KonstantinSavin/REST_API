package apiserver

import "effective-mobile/music-lib/internal/storage"

type Config struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
	Storage  *storage.Config
}

func NewCfg() *Config {
	return &Config{
		Addr:     ":8080",
		LogLevel: "debug",
		Storage:  storage.NewCfg(),
	}
}
