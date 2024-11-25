package apiserver

import (
	"net/http"

	_ "effective-mobile/music-lib/docs"
	"effective-mobile/music-lib/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type serv interface {
	AddSong(song model.Song) (*model.EnrichedSong, error)
	DeleteSong(id string) error
	UpdateSong(id string, s *model.Song) (*model.Song, error)
	GetSongs(f *model.Filter) ([]*model.Song, bool, error)
}

type server struct {
	router *gin.Engine
	logger *logrus.Logger
	// storage *storage.Storage
	service serv
}

func newServer(logger *logrus.Logger, service serv) *server {
	srv := &server{
		router:  gin.New(),
		logger:  logger,
		service: service,
	}

	srv.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	srv.router.POST("/songs", srv.handlerGetSongs)
}
