package sqldb

import (
	"effective-mobile/music-lib/internal/model"
	"effective-mobile/music-lib/internal/storage"
	"fmt"

	"github.com/sirupsen/logrus"
)

type SongRep struct {
	storage *Storage
	logger  *logrus.Logger
}

func (r *SongRep) CreateSong(s *model.Song) error {
	q := `INSERT INTO songs (song_name,group_name) VALUES ($1, $2) RETURNING id`

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", q))

	if err := r.storage.db.QueryRow(
		q,
		s.Name,
		s.Group,
	).Scan(&s.ID); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return err
	}
	return nil
}

func (r *SongRep) DeleteSong(id string) error {
	q := `DELETE FROM songs WHERE id = $1`

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", q))

	if _, err := r.storage.db.Exec(q, id); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return err
	}
	return nil
}

func (r *SongRep) UpdateSong(id string, s *model.Song) (*model.Song, error) {
	q := `UPDATE songs
		SET
			song_name = COALESCE($2, song_name),
			group_name = COALESCE($3, group_name)
		WHERE id = $1
		RETURNING id, song_name, group_name`

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", q))

	song := &model.Song{}
	if err := r.storage.db.QueryRow(
		q,
		id,
		s.Name,
		s.Group,
	).Scan(&song.ID, &song.Name, &song.Group); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return nil, err
	}

	return song, nil
}

func (r *SongRep) GetSongs(f *storage.Filter) ([]*model.Song, bool, error) {
	var hasNextPage bool = false

	songs, err := r.FilterSongs(*f)
	if err != nil {
		return nil, false, err
	}

	if len(songs) > *f.PerPage {
		songs = songs[:len(songs)-1]
		hasNextPage = true
	}

	return songs, hasNextPage, nil
}

// TODO
func (r *SongRep) GetText(id string) (string, error) {
	s := ""
	if err := r.storage.db.QueryRow(
		"SELECT id, song_name, group_name FROM songs WHERE id = $1",
		id,
	).Scan(); err != nil {
		return s, err
	}

	return s, nil
}
