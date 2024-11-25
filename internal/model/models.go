package model

type Song struct {
	SongID int    `json:"id"`
	Name   string `json:"song"`
	Group  string `json:"group"`
}

type EnrichedSong struct {
	SongID      int    `json:"id"`
	Name        string `json:"song"`
	Group       string `json:"group"`
	GroupID     int    `json:"groupID"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
