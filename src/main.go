package main

import (
	"autoposter/server"
	"autoposter/config"
)

func main() {
	config.InitEnv()
	server.Run()
}