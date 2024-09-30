package storage_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=effmob password=effmob host=localhost port=5432 dbname=music_lib_db sslmode=disable"
	}

	os.Exit(m.Run())
}
