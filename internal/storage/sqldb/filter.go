package sqldb

import (
	"database/sql"
	"effective-mobile/music-lib/internal/model"
	"effective-mobile/music-lib/internal/storage"
	"fmt"
	"strconv"
)

func (r SongRep) FilterSongs(f storage.Filter) ([]*model.Song, error) {
	rows, err := r.queryWithFilter(f)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*model.Song
	for rows.Next() {
		s := new(model.Song)
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Group,
		)
		if err != nil {
			return nil, err
		}
		songs = append(songs, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if songs == nil {
		return []*model.Song{}, nil
	}

	return songs, err
}

func (r SongRep) queryWithFilter(f storage.Filter) (*sql.Rows, error) {
	prefix := `
	SELECT id, song_name, group_name
	FROM songs
`

	query := ` WHERE`
	var filterValues []interface{}

	if f.Name != nil {
		filterValues = append(filterValues, "%"+*f.Name+"%")
		if query != ` WHERE` {
			query += ` AND`
		}
		query += ` song_name LIKE $` + strconv.Itoa(len(filterValues))

	}

	if f.Group != nil {
		filterValues = append(filterValues, "%"+*f.Group+"%")
		if query != ` WHERE` {
			query += ` AND`
		}
		query += ` group_name LIKE $` + strconv.Itoa(len(filterValues))

	}

	if query == ` WHERE` {
		query += ` 1=1`
	}

	query += ` ORDER BY id `

	limit := *f.PerPage + 1
	filterValues = append(filterValues, limit)
	query += ` LIMIT $` + strconv.Itoa(len(filterValues)) // TODO

	offset := (*f.Page - 1) * *f.PerPage

	filterValues = append(filterValues, offset)
	query += ` OFFSET $` + strconv.Itoa(len(filterValues))

	sql := prefix + query
	fmt.Println(sql)
	fmt.Println(filterValues...)
	return r.storage.songRep.storage.db.Query(sql, filterValues...)
}
