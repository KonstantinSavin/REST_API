package sqldb

import (
	"database/sql"
	"effective-mobile/music-lib/internal/model"
	"fmt"
	"strconv"
)

func (r SongRep) FilterSongs(f model.Filter) ([]*model.EnrichedSong, error) {
	r.logger.Debugf("SongRep FilterSongs")

	rows, err := r.queryWithFilter(f)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*model.EnrichedSong
	for rows.Next() {
		s := new(model.EnrichedSong)
		err := rows.Scan(
			&s.SongID,
			&s.Name,
			&s.GroupID,
			&s.ReleaseDate,
			&s.Text,
			&s.Link,
		)
		if err != nil {
			return nil, err
		}

		group, err := r.getGroupNameByID(s.GroupID)
		if err != nil {
			return nil, err
		}

		s.Group = group
		songs = append(songs, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if songs == nil {
		return []*model.EnrichedSong{}, nil
	}

	return songs, err
}

func (r SongRep) getGroupNameByID(id int) (string, error) {
	r.logger.Debugf("SongRep getGroupNameByID")

	q := `SELECT group_name 
		FROM group_names 
		WHERE id = $1;`

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", q))

	var groupName string

	if err := r.storage.db.QueryRow(
		q,
		id,
	).Scan(&groupName); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return "", err
	}

	return groupName, nil
}

func (r SongRep) queryWithFilter(f model.Filter) (*sql.Rows, error) {
	r.logger.Debugf("SongRep queryWithFilter")

	if f.Group != nil {
		qg := `SELECT id 
		FROM group_names 
		WHERE group_name = $1;`

		r.logger.Debugf(fmt.Sprintf("SQL Query: %s, $1 = %s", qg, *f.Group))

		if err := r.storage.db.QueryRow(
			qg,
			f.Group,
		).Scan(&f.GroupID); err != nil {
			r.logger.Errorf("Ошибка SQL: %s", err)
			return nil, err
		}
	}

	prefix := `
	SELECT id, song_name, group_id, release_date, text, link
	FROM songs
`

	query := ` WHERE`
	var filterValues []interface{}

	if f.Name != nil {
		filterValues = append(filterValues, *f.Name)
		if query != ` WHERE` {
			query += ` AND`
		}
		query += ` song_name LIKE $` + strconv.Itoa(len(filterValues))

	}

	if f.GroupID != nil {
		filterValues = append(filterValues, *f.GroupID)
		if query != ` WHERE` {
			query += ` AND`
		}
		query += ` group_id = $` + strconv.Itoa(len(filterValues))

	}

	if query == ` WHERE` {
		query += ` 1=1`
	}

	query += ` ORDER BY id `

	limit := *f.PerPage + 1
	filterValues = append(filterValues, limit)
	query += ` LIMIT $` + strconv.Itoa(len(filterValues))

	offset := (*f.Page - 1) * *f.PerPage

	filterValues = append(filterValues, offset)
	query += ` OFFSET $` + strconv.Itoa(len(filterValues))

	sql := prefix + query

	r.logger.Debugf(fmt.Sprintf("SQL Query: %s", sql))

	return r.storage.db.Query(sql, filterValues...)
}
