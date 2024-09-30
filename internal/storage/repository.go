package storage

import "effective-mobile/music-lib/internal/model"

type SongRep interface {
	CreateSong(s *model.Song) error
	DeleteSong(id string) error
	UpdateSong(id string, s *model.Song) (*model.Song, error)
	GetSongs(s *model.Song) ([]*model.Song, bool, error)
	GetText(id string) (string, error)
}
