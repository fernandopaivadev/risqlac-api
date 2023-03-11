package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"risqlac-api/environment"
	"risqlac-api/types"
)

var Instance *gorm.DB

func autoMigrateModel(db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(model)
	if err != nil {
		panic("Error migrating model => " + err.Error())
	}
}

func Connect() {
	db, err := gorm.Open(sqlite.Open(environment.Get().DATABASE_FILE), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	autoMigrateModel(db, &types.Product{})
	autoMigrateModel(db, &types.User{})

	Instance = db
}
