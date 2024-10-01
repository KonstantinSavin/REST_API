package apiserver

import (
	"database/sql"
	"effective-mobile/music-lib/internal/storage/sqldb"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Start(cfg *Config, logger *logrus.Logger) error {
	logger.Debugf("подключаем базу данных по адресу: %s", cfg.DatabaseURL)
	db, err := newDB(cfg.DatabaseURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("база данных подключена")

	defer db.Close()
	storage := sqldb.New(db)
	srv := newServer(storage)

	logger.Debugf("подключаем сервер по адресу %s", cfg.Addr)
	return http.ListenAndServe(cfg.Addr, srv)
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
