package storage

import (
	"effective-mobile/music-lib/internal/model"
)

type SongRep interface {
	CreateSong(s *model.EnrichedSong) (*model.EnrichedSong, error)
	DeleteSong(id string) error
	UpdateSong(id string, s *model.Song) (*model.Song, error)
	GetSongs(f *model.Filter) ([]*model.Song, bool, error)
}
