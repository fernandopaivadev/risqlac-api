package infra

import (
	"errors"
	"risqlac-api/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	Instance *gorm.DB
}

var Database database

func (database *database) Connect() error {
	db, err := gorm.Open(sqlite.Open(Environment.Variables.DatabaseFile), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect to the database")
	}

	err = db.AutoMigrate(&types.User{})
	if err != nil {
		return errors.New("Error migrating model => " + err.Error())
	}

	err = db.AutoMigrate(&types.Product{})
	if err != nil {
		return errors.New("Error migrating model => " + err.Error())
	}

	database.Instance = db

	return nil
}
