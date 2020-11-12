package server

import (
	"autoposter/config"
	"autoposter/routes"
	"fmt"
	"log"
	"net/http"
)

// Server - ...
type Server struct {
	httpServer *http.Server
}

// Run - running server
func (s *Server) Run() error {
	r := routes.NewRouter()
	cfg := config.GetConfig()

	fmt.Printf("Server is running on :%s\n", cfg.Port)

	err := http.ListenAndServe(cfg.Port, r)
	if err != nil {
		log.Fatal(err)
	}
}
