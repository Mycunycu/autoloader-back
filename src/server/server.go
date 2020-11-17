package server

import (
	"fmt"
	"log"
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

	fmt.Printf("Server is running on :%s\n", port)

	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
