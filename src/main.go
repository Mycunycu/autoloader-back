package main

import (
	"autoposter/config"
	"autoposter/mongodb"
	"autoposter/routes"
	"autoposter/server"
)

func main() {
	config.InitEnv()
	cfg := config.GetConfig()

	mongoStore := new(mongodb.MongoDataStore)
	mongoStore.CreateDataStore(cfg.DbName, cfg.MongoDbURL)

	srv := new(server.Server)
	router := new(routes.Router)
	srv.Run(cfg.Port, router.InitRouter())
}
