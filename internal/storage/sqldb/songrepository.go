package sqldb

import (
	"effective-mobile/music-lib/internal/model"
	"fmt"

	"github.com/sirupsen/logrus"
)

type SongRep struct {
	storage *Storage
	logger  *logrus.Logger
}

func (r *SongRep) CreateSong(s *model.EnrichedSong) (*model.EnrichedSong, error) {
	r.logger.Debugf("SongRep CreateSong")

	qg := `WITH inserted AS (
    	INSERT INTO group_names (group_name) VALUES ($1) 
    	ON CONFLICT (group_name) DO NOTHING
    	RETURNING id
		)
		SELECT id FROM inserted
		UNION ALL
		SELECT id FROM group_names WHERE group_name = $1 LIMIT 1;`

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", qg))

	if err := r.storage.db.QueryRow(
		qg,
		s.Group,
	).Scan(&s.GroupID); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return s, err
	}

	qs := `INSERT INTO songs (song_name, group_id, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", qs))

	if err := r.storage.db.QueryRow(
		qs,
		s.Name,
		s.GroupID,
		s.ReleaseDate,
		s.Text,
		s.Link,
	).Scan(&s.SongID); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return s, err
	}
	return s, nil
}

func (r *SongRep) DeleteSong(id string) error {
	r.logger.Debugf("SongRep DeleteSong")
	q := `DELETE FROM songs WHERE id = $1`

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", q))

	if _, err := r.storage.db.Exec(q, id); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return err
	}
	return nil
}

func (r *SongRep) UpdateSong(id string, s *model.Song) (*model.Song, error) {
	r.logger.Debugf("SongRep UpdateSong")

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
	).Scan(&song.SongID, &song.Name, &song.Group); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return nil, err
	}

	return song, nil
}

func (r *SongRep) GetSongs(f *model.Filter) ([]*model.Song, bool, error) {
	r.logger.Debugf("SongRep GetSongs")

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
