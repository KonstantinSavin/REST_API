package apiserver

import (
	"effective-mobile/music-lib/internal/storage"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	cfg     *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *APIServer {
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	server.logger.Info("запуск сервера")

	if err := server.configureLogger(); err != nil {
		return err
	}

	server.configureRouter()

	if err := server.configureStorage(); err != nil {
		return err
	}

	server.logger.Info("api server начал работу")

	return http.ListenAndServe(server.cfg.Addr, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.cfg.LogLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)
	server.logger.Infof("уровень логирования %s", level)

	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) configureStorage() error {
	st := storage.New(server.cfg.Storage)
	if err := st.Open(); err != nil {
		return err
	}

	server.storage = st
	server.logger.Info("база данных подключена")

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}
