package server

import (
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func (s *Server) InitServer(port string, handler http.Handler) error {
	s.Server = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.Server.ListenAndServe()
}
