package main

import (
	"effective-mobile/music-lib/internal/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := apiserver.NewCfg()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
