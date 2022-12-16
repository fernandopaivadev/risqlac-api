package database

import (
	"os"
	"risqlac-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Users []models.User
var Products []models.Product

var Instance *gorm.DB

func Setup() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_FILE")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})

	Instance = db
}
