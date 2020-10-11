package main

import (
	"autoposter/services/mongodb"
	"autoposter/server"
	"autoposter/config"
)

func main() {
	config.InitEnv()
	mongodb.CreateDataStore()
	server.Run()
}
