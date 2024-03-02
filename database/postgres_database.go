package database

import (
	"awesomeProject/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(config *config.Config) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Db.Host,
		config.Db.User,
		config.Db.Password,
		config.Db.DBName,
		config.Db.Port,
		config.Db.SSLMode,
		config.Db.TimeZone,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{Db: database}
}

func (postgres postgresDatabase) GetDatabase() *gorm.DB {
	return postgres.Db
}
