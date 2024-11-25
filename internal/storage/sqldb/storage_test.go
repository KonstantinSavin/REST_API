package sqldb_test

import (
	"os"
	"testing"
)

var (
	testDatabaseURL string
)

func TestMain(m *testing.M) {
	testDatabaseURL = os.Getenv("DB_URL_TEST")
	if testDatabaseURL == "" {
		testDatabaseURL = "user=effmob password=effmob host=localhost port=15432 dbname=ml_db sslmode=disable"
	}

	os.Exit(m.Run())
}
