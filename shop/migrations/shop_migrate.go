package main

import (
	"awesomeProject/config"
	"awesomeProject/database"
	"awesomeProject/shop/entities"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	shopDatabase(db)
}

func shopDatabase(db database.Database) {
	db.GetDatabase().Migrator().CreateTable(&entities.Shop{})
	db.GetDatabase().CreateInBatches([]entities.Shop{
		{Amount: 1},
		{Amount: 2},
		{Amount: 2},
		{Amount: 5},
		{Amount: 3},
	}, 10)
}
