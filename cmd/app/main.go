package main

import (
	_ "effective-mobile/music-lib/docs"
	"effective-mobile/music-lib/internal/apiserver"
	"effective-mobile/music-lib/internal/config"
	"effective-mobile/music-lib/pkg/logging"

	"github.com/joho/godotenv"
)

// Swagger
//
// @title        Music library
// @version      1.0
// @description  API server for music library
// @host         localhost:8000
// @BasePath     /
// @schemes      http

func main() {
	logger := logging.GetLogger()
	logger.Info("запуск приложения")

	logger.Debug("парсим конфиг")
	if err := godotenv.Load(); err != nil {
		logger.Info("файл .env не найден")
	}

	cfg := config.GetConfig()
	logger.Infof("конфиг получен: port: %s, DB_url: %s, API_URL: %s",
		cfg.Port, cfg.DBURL, cfg.APIURL)

	logger.Debug("запускаем сервер")
	if err := apiserver.Start(cfg, logger); err != nil {
		logger.Fatal(err)
	}
}
