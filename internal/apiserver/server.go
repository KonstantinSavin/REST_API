package apiserver

import (
	"effective-mobile/music-lib/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	storage storage.Storage
}

func newServer(storage storage.Storage) *server {
	srv := &server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		storage: storage,
	}

	srv.configureRouter()

	return srv
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}

func (srv *server) configureRouter() {
	srv.router.HandleFunc("/songs", srv.handleSongsCreate()).Methods("POST")
}

func (srv *server) handleSongsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
