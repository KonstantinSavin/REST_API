package apiserver

type Config struct {
	Addr        string `toml:"addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"db_url"`
}

func NewCfg() *Config {
	return &Config{
		Addr:     ":8080",
		LogLevel: "debug",
	}
}
