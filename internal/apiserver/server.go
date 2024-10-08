package apiserver

import (
	"effective-mobile/music-lib/internal/storage"
	"net/http"

	_ "effective-mobile/music-lib/docs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
