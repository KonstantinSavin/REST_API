package apiserver

import (
	"effective-mobile/music-lib/internal/storage/sqldb"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const dbURL = "user=effmob password=effmob host=localhost port=5432 dbname=ml_db sslmode=disable"

func TestServer(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/songs", nil)
	db, err := newDB(dbURL)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	storage := sqldb.New(db)
	srv := newServer(logrus.New(), storage)
	srv.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusBadRequest)
}
