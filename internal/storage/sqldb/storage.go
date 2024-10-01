package sqldb

import (
	"database/sql"
	"effective-mobile/music-lib/internal/storage"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	db      *sql.DB
	songRep *SongRep
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (st *Storage) Song() storage.SongRep {
	if st.songRep != nil {
		return st.songRep
	}

	st.songRep = &SongRep{
		storage: st,
		logger:  logrus.New(),
	}

	return st.songRep
}
