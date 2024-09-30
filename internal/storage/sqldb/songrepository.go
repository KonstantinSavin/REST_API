package sqldb

import "effective-mobile/music-lib/internal/model"

type SongRep struct {
	storage *Storage
}

func (r *SongRep) CreateSong(s *model.Song) error {
	if err := r.storage.db.QueryRow(
		"INSERT INTO songs (song_name,group_name) VALUES ($1, $2) RETURNING id",
		s.Name,
		s.Group,
	).Scan(&s.ID); err != nil {
		return err
	}
	return nil
}

func (r *SongRep) DeleteSong(id string) error {
	if _, err := r.storage.db.Exec(
		"DELETE FROM songs WHERE id = $1",
		id,
	); err != nil {
		return err
	}
	return nil
}

func (r *SongRep) UpdateSong(id string, s *model.Song) (*model.Song, error) {
	song := &model.Song{}
	if err := r.storage.db.QueryRow(
		`UPDATE songs
		SET
			song_name = COALESCE($2, song_name),
			group_name = COALESCE($3, group_name)
		WHERE id = $1
		RETURNING id, song_name, group_name`,
		id,
		s.Name,
		s.Group,
	).Scan(&song.ID, &song.Name, &song.Group); err != nil {
		return nil, err
	}

	return song, nil
}

//TODO
func (r *SongRep) GetSongs(s *model.Song) ([]*model.Song, bool, error) {
	return nil, false, nil
}

//TODO
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