package main

import (
	"awesomeProject/config"
	"awesomeProject/database"
	"awesomeProject/server"
)

//	@title			Swaggerbleat nahuy
//	@version		1.0
//	@description	omaigadebat
func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDatabase()).Start()
}
