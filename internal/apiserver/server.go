package apiserver

import (
	"effective-mobile/music-lib/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type server struct {
	router  *gin.Engine
	logger  *logrus.Logger
	storage storage.Storage
}

func newServer(logger *logrus.Logger, storage storage.Storage) *server {
	srv := &server{
		router:  gin.New(),
		logger:  logger,
		storage: storage,
	}

	srv.configureRouter()

	return srv
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}

func (srv *server) configureRouter() {
	srv.router.POST("/add", srv.handlerAddSong)
	srv.router.DELETE("/delete/:id", srv.handlerDeleteSong)
	srv.router.PATCH("/update/:id", srv.handlerUpdateSong)
}
