package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	config  *Config
	db      *sql.DB
	songRep *SongRep
}

func New(cfg *Config) *Storage {
	return &Storage{
		config: cfg,
	}
}

func (st *Storage) Open() error {
	dsn := st.config.DBURL
	if dsn == "" {
		dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			st.config.Username, st.config.Password, st.config.Host, st.config.Port, st.config.Database)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	st.db = db

	return nil
}

func (st *Storage) Close() {
	st.db.Close()
}

func (st *Storage) Song() *SongRep {
	if st.songRep != nil {
		return st.songRep
	}

	st.songRep = &SongRep{
		storage: st,
	}

	return st.songRep
}
