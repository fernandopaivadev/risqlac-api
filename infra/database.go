package infra

import (
	"errors"
	"risqlac-api/application/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	Instance *gorm.DB
}

var Database database

func (database *database) Connect() error {
	dsn := Environment.Variables.DatabaseUrl + "?tls=true&parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect to the database")
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		return errors.New("Error migrating model => " + err.Error())
	}

	err = db.AutoMigrate(&models.Product{})

	if err != nil {
		return errors.New("Error migrating model => " + err.Error())
	}

	database.Instance = db

	return nil
}
