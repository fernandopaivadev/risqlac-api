package database

import (
	"risqlac-api/environment"
	"risqlac-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(environment.Get().DATABASE_FILE), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})

	Instance = db
}
