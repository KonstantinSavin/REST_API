package service

import (
	"effective-mobile/music-lib/internal/model"
	"effective-mobile/music-lib/internal/storage"

	"github.com/sirupsen/logrus"
)

type Service struct {
	storage storage.Storage
	apiUrl  string
	logger  *logrus.Logger
}

func NewService(storage storage.Storage, apiUrl string, logger *logrus.Logger) *Service {
	return &Service{
		storage: storage,
		apiUrl:  apiUrl,
		logger:  logger,
	}
}

func (s *Service) AddSong(song model.Song) (*model.EnrichedSong, error) {
	s.logger.Debug("Service AddSong")

	info, err := s.fetchInfoFromAPI(song.Group, song.Name)
	if err != nil {
		return &model.EnrichedSong{}, err
	}

	enrichedSong := &model.EnrichedSong{
		SongID:      song.SongID,
		Name:        song.Name,
		Group:       song.Group,
		ReleaseDate: info.ReleaseDate,
		Text:        info.Text,
		Link:        info.Link,
	}

	return s.storage.Song().CreateSong(enrichedSong)
}

func (s *Service) DeleteSong(id string) error {
	s.logger.Debug("Service DeleteSong")
	return s.storage.Song().DeleteSong(id)
}

func (s *Service) UpdateSong(id string, newSong *model.EnrichedSong) (*model.EnrichedSong, error) {
	s.logger.Debug("Service UpdateSong")
	return s.storage.Song().UpdateSong(id, newSong)
}

func (s *Service) GetSongs(f *model.Filter) ([]*model.EnrichedSong, bool, error) {
	s.logger.Debug("Service GetSongs")
	return s.storage.Song().GetSongs(f)
}

func (s *Service) GetCouplets(stp *model.SongTextPagination) (*model.PaginatedText, bool, error) {
	s.logger.Debug("Service GetCouplets")
	return s.storage.Song().GetCouplets(stp)
}
