package service

import (
	"effective-mobile/music-lib/internal/model"
	"fmt"
	"strconv"
)

func (f Filter) Update() Filter {
	return Filter{
		Pages: 10,
		ID:    f.ID,
		Name:  f.Name,
		Group: f.Group,
	}
}

func (f Filter) FilterSongs() ([]*model.Song, error) {
	prefix := `
	SELECT id, song_name, group_name
	FROM songs
`

	query := ` WHERE`
	var filterValues []interface{}

	if f.Name != "" {
		filterValues = append(filterValues, "%"+f.Name+"%")
		if query != ` WHERE` {
			query += ` AND`
		}
		query += ` song_name LIKE $` + strconv.Itoa(len(filterValues))
		fmt.Println(query, filterValues)

	}

	if f.Group != "" {
		filterValues = append(filterValues, "%"+f.Name+"%")
		if query != ` WHERE` {
			query += ` AND`
		}
		query += ` group_name LIKE $` + strconv.Itoa(len(filterValues))
		fmt.Println(query, filterValues)

	}

	if query == ` WHERE` {
		query += ` 1=1`
	}

	query += ` ORDER BY id `

	limit := f.Pages + 1
	filterValues = append(filterValues, limit)
	query += ` LIMIT $` + strconv.Itoa(len(filterValues))

	sql := prefix + query
	fmt.Println(sql)

	return nil, nil
}
