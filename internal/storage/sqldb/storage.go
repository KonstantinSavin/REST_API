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
	logger  *logrus.Logger
}

func New(db *sql.DB, logger *logrus.Logger) *Storage {
	return &Storage{
		db:     db,
		logger: logger,
	}
}

func (st *Storage) Song() storage.SongRep {
	if st.songRep != nil {
		return st.songRep
	}

	st.songRep = &SongRep{
		storage: st,
		logger:  st.logger,
	}

	return st.songRep
}
