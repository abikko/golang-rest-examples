package main

import (
	"awesomeProject/config"
	"awesomeProject/database"
	"awesomeProject/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDatabase()).Start()
}
