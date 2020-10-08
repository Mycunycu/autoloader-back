package server

import (
	"net/http"
	"autoposter/routes"
	"autoposter/config"
	"fmt"
	"log"
)

// Run - running server
func Run() {
	r := routes.NewRouter()
	cfg := config.GetConfig()
	
	go log.Fatal(http.ListenAndServe(cfg.Port, r))

	fmt.Printf("Server is running on :%s\n", cfg.Port)
}