package database

import (
	"djaeger-bot/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Init() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")

	db, connErr := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if connErr != nil {
		log.Fatalln(connErr)
	}

	err := db.AutoMigrate(&models.Activity{})
	if err != nil {
		return nil
	}

	return db
}
