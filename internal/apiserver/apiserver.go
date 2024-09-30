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

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStorage(); err != nil {
		return err
	}

	s.logger.Info("api server начал работу")

	return http.ListenAndServe(s.cfg.Addr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.cfg.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStorage() error {
	st := storage.New(s.cfg.Storage)
	if err := st.Open(); err != nil {
		return err
	}

	s.storage = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}
