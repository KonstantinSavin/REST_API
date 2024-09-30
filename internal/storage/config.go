package storage

type Config struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Database string `toml:"dbname"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBURL    string `toml:"db_url"`
}

func NewCfg() *Config {
	return &Config{}
}
