package storage

import (
	"effective-mobile/music-lib/internal/model"
)

type SongRep interface {
	CreateSong(s *model.EnrichedSong) (*model.EnrichedSong, error)
	DeleteSong(id string) error
	UpdateSong(id string, s *model.EnrichedSong) (*model.EnrichedSong, error)
	GetSongs(f *model.Filter) ([]*model.EnrichedSong, bool, error)
}
