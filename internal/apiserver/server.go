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

// func (srv *server) handleSongsCreate() http.HandlerFunc {

// 	type request struct {
// 		Name  string `json:"name"`
// 		Group string `json:"group"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		req := &request{}
// 		srv.logger.Infof("запрос %s", r)
// 		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 			srv.logger.Error(err, req)
// 			srv.error(w, r, http.StatusBadRequest, err)
// 			return
// 		}

// 		s := &model.Song{
// 			Name:  req.Name,
// 			Group: req.Group,
// 		}
// 		if err := srv.storage.Song().CreateSong(s); err != nil {
// 			srv.error(w, r, http.StatusUnprocessableEntity, err)
// 			return
// 		}

// 		srv.respond(w, r, http.StatusCreated, s)
// 	}
// }

// func (srv *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
// 	srv.respond(w, r, code, map[string]string{"error": err.Error()})
// }

// func (srv *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
// 	w.WriteHeader(code)
// 	if data != nil {
// 		json.NewEncoder(w).Encode(data)
// 	}
// }
