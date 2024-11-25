package apiserver

import (
	"effective-mobile/music-lib/internal/service"
	"effective-mobile/music-lib/internal/storage/sqldb"
	"effective-mobile/music-lib/pkg/logging"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const dbURL = "user=effmob password=effmob host=localhost port=15432 dbname=ml_db sslmode=disable"
const apiurl = ""

func TestServer(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/songs", nil)
	db, err := newDB(dbURL)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	logger := logging.GetLogger()
	storage := sqldb.New(db, logger)
	srv := newServer(logrus.New(), service.NewService(storage, apiurl, logger))
	srv.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusBadRequest)
}
