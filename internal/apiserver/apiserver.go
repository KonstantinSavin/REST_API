package apiserver

import (
	"database/sql"
	"effective-mobile/music-lib/internal/storage/sqldb"
	"fmt"
	"net/http"
)

func Start(cfg *Config) error {
	fmt.Println(cfg.DatabaseURL)
	db, err := newDB(cfg.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	storage := sqldb.New(db)
	srv := newServer(storage)

	return http.ListenAndServe(cfg.Addr, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
