package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Server - ...
type Server struct {
	httpServer *http.Server
}

// Run - ...
func (s *Server) Run(port string, handler http.Handler) {
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logrus.Printf("Server is running on :%s\n", port)

	err := s.httpServer.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
