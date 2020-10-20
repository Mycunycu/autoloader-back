package main

import (
	"autoposter/config"
	"autoposter/server"
	"autoposter/services/mongodb"
)

func main() {
	config.InitEnv()
	mongodb.CreateDataStore()
	server.Run()
}
