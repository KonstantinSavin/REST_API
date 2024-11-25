package sqldb_test

import (
	"effective-mobile/music-lib/internal/model"
	"effective-mobile/music-lib/internal/storage/sqldb"
	"effective-mobile/music-lib/pkg/logging"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSongRepository_CreateSong(t *testing.T) {
	db, teardown := sqldb.TestDB(t, testDatabaseURL)
	defer teardown("songs")

	st := sqldb.New(db, logging.GetLogger())

	s := model.EnrichedSong{
		Name:        "Gimme!",
		Group:       "ABBA",
		ReleaseDate: "",
		Text:        "",
		Link:        "",
	}
	newSong, err := st.Song().CreateSong(&s)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, newSong.Name, s.Name)
	assert.Equal(t, newSong.Group, s.Group)
}

// func TestSongRepository_DeleteSong(t *testing.T) {
// 	db, teardown := sqldb.TestDB(t, databaseURL)
// 	defer teardown("songs")

// 	st := sqldb.New(db)

// 	id := ""
// 	err := st.Song().DeleteSong(id)
// 	assert.Error(t, err)

// 	s := &model.Song{
// 		Name:  "Gimme!",
// 		Group: "ABBA",
// 	}
// 	err = st.Song().CreateSong(s)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, s.SongID)

// 	err = st.Song().DeleteSong(s.SongID)
// 	assert.NoError(t, err)
// }

// func TestSongRepository_UpdateSong(t *testing.T) {
// 	db, teardown := sqldb.TestDB(t, databaseURL)
// 	defer teardown("songs")

// 	st := sqldb.New(db)

// 	s := &model.Song{
// 		Name:  "Gimme!",
// 		Group: "ABBA",
// 	}
// 	err := st.Song().CreateSong(s)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, s)
// 	assert.Equal(t, "Gimme!", s.Name)
// 	assert.Equal(t, "ABBA", s.Group)

// 	s, err = st.Song().UpdateSong(s.SongID, &model.Song{
// 		Name:  "Can't Stop",
// 		Group: "RHCP",
// 	})
// 	assert.NoError(t, err)
// 	assert.NotNil(t, s)
// 	assert.Equal(t, "Can't Stop", s.Name)
// 	assert.Equal(t, "RHCP", s.Group)
// }

// TODO
// func TestSongRepository_GetSong(t *testing.T) {
// 	db, teardown := storage.TestDB(t, databaseURL)
// 	defer teardown("songs")

// st := storage.New(db)
// }

// TODO
// func TestSongRepository_GetText(t *testing.T) {
// 	db, teardown := storage.TestDB(t, databaseURL)
// 	defer teardown("songs")

// st := storage.New(db)
// }
