package sqldb

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDB(t *testing.T, testDatabaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", testDatabaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
