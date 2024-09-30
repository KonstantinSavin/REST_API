package storage_test

import (
	"effective-mobile/music-lib/internal/model"
	"effective-mobile/music-lib/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSongRepository_CreateSong(t *testing.T) {
	st, teardown := storage.TestStorage(t, databaseURL)
	defer teardown("songs")

	s, err := st.Song().CreateSong(&model.Song{
		Name:  "Gimme!",
		Group: "ABBA",
	})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "Gimme!", s.Name)
	assert.Equal(t, "ABBA", s.Group)
}

func TestSongRepository_DeleteSong(t *testing.T) {
	st, teardown := storage.TestStorage(t, databaseURL)
	defer teardown("songs")

	id := ""
	err := st.Song().DeleteSong(id)
	assert.Error(t, err)

	s, err := st.Song().CreateSong(&model.Song{
		Name:  "Can't Stop",
		Group: "RHCP",
	})
	assert.NoError(t, err)
	assert.NotNil(t, s.ID)

	err = st.Song().DeleteSong(s.ID)
	assert.NoError(t, err)
}

func TestSongRepository_UpdateSong(t *testing.T) {
	st, teardown := storage.TestStorage(t, databaseURL)
	defer teardown("songs")

	s, err := st.Song().CreateSong(&model.Song{
		Name:  "Gimme!",
		Group: "ABBA",
	})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "Gimme!", s.Name)
	assert.Equal(t, "ABBA", s.Group)

	s, err = st.Song().UpdateSong(s.ID, &model.Song{
		Name:  "Can't Stop",
		Group: "RHCP",
	})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "Can't Stop", s.Name)
	assert.Equal(t, "RHCP", s.Group)
}

// TODO
// func TestSongRepository_GetSong(t *testing.T) {
// 	st, teardown := storage.TestStorage(t, databaseURL)
// 	defer teardown("songs")
// }

// TODO
// func TestSongRepository_GetText(t *testing.T) {
// 	st, teardown := storage.TestStorage(t, databaseURL)
// 	defer teardown("songs")
// }
