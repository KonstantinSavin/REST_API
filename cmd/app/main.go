package main

import (
	"effective-mobile/music-lib/internal/apiserver"
	"effective-mobile/music-lib/pkg/logging"
	"flag"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	logger := logging.GetLogger()
	logger.Info("запуск приложения")

	logger.Debug("парсим конфиг")
	flag.Parse()

	cfg := apiserver.NewCfg()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("конфиг получен")

	logger.Debug("запускаем сервер")
	if err := apiserver.Start(cfg, logger); err != nil {
		logger.Fatal(err)
	}
}
