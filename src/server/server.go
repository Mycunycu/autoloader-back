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

	fmt.Printf("Server is running on :%s\n", cfg.Port)

	err := http.ListenAndServe(cfg.Port, r)

	if err != nil {
		log.Fatal(err)
	}
}