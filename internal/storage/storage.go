package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	db     *sql.DB
}

func New(cfg *Config) *Storage {
	return &Storage{
		config: cfg,
	}
}

// "host=localhost dbname=music_lib sslmode=disable"

func (s *Storage) Open() error {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s "+
		"dbname=%s sslmode=disable", s.config.Username, s.config.Password, s.config.Host, s.config.Port, s.config.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}
