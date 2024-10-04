package apiserver

import (
	"database/sql"
	"effective-mobile/music-lib/internal/config"
	"effective-mobile/music-lib/internal/storage/sqldb"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Start(cfg *config.Config, logger *logrus.Logger) error {
	logger.Debugf("подключаем базу данных по адресу: %s", cfg.DBURL)
	db, err := newDB(cfg.DBURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("база данных подключена")

	if err := sqldb.MigrationsUp(db); err != nil {
		return err
	}
	logger.Info("db мигрировало")

	defer db.Close()
	storage := sqldb.New(db)
	srv := newServer(logger, storage)

	logger.Debugf("подключаем сервер по адресу %s", cfg.Port)
	logger.Info("приложение запущено")
	return http.ListenAndServe(cfg.Port, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
