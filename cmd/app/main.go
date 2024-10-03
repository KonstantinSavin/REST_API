package main

import (
	// "effective-mobile/music-lib/internal/apiserver"
	"effective-mobile/music-lib/internal/apiserver"
	"effective-mobile/music-lib/internal/config"
	"effective-mobile/music-lib/pkg/logging"

	"github.com/joho/godotenv"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("запуск приложения")

	logger.Debug("парсим конфиг")
	if err := godotenv.Load(); err != nil {
		logger.Info("файл .env не найден")
	}

	cfg := config.GetConfig()
	logger.Infof(`конфиг получен:
	port: %s
	loglevel: %s
	DB url: %s`, cfg.Port, cfg.LogLevel, cfg.DBURL)

	logger.Debug("запускаем сервер")
	if err := apiserver.Start(cfg, logger); err != nil {
		logger.Fatal(err)
	}
}
