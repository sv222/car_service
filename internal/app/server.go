package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Router     *mux.Router
	Logger     *logrus.Logger
	Config     *Config
	HTTPServer *http.Server
}

func NewServer() *Server {
	r := mux.NewRouter()
	c := NewConfig()
	s := &http.Server{
		Addr:              ":" + c.Port,
		Handler:           r,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 15,
		MaxHeaderBytes:    1 << 20, // 1 MB
	}

	return &Server{
		Router:     r,
		Logger:     logrus.New(),
		Config:     c,
		HTTPServer: s,
	}
}

func (s *Server) Start() error {
	s.Logger.Println("staring server starting on port %d", s.Config.Port)
	return s.HTTPServer.ListenAndServe()
}
