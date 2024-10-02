package service

import "effective-mobile/music-lib/internal/model"

type Filter struct {
	Pages int    `json:"pages,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"song,omitempty"`
	Group string `json:"group,omitempty"`
}

type FilteredSongs struct {
	Songs []*model.Song `json:"songs"`
}
