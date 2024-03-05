package main

import (
	"awesomeProject/config"
	"awesomeProject/database"
	"awesomeProject/server"
)

//	@title			Swagger
//	@version		1.0
//	@description	Swagger docs example for pet project
func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDatabase()).Start()
}
