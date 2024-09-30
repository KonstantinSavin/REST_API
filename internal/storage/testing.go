package storage

import (
	"fmt"
	"strings"
	"testing"
)

func TestStorage(t *testing.T, databaseURL string) (*Storage, func(...string)) {
	t.Helper()

	cfg := NewCfg()
	cfg.DBURL = databaseURL
	st := New(cfg)
	if err := st.Open(); err != nil {
		t.Fatal(err)
	}

	return st, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := st.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ","))); err != nil {
				t.Fatal(err)
			}
		}

		st.Close()
	}
}
