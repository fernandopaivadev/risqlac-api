package models

import (
	"time"
)

type User struct {
	Id         uint64    `json:"id" gorm:"unique; autoIncrement; primaryKey; <-:create"`
	Username   string    `json:"username" gorm:"unique" validate:"required"`
	Email      string    `json:"email" gorm:"unique" validate:"required,email"`
	Name       string    `json:"name" validate:"required"`
	Phone      string    `json:"phone" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Is_admin   bool      `json:"is_admin"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Product struct {
	Id              uint64    `json:"id" gorm:"autoIncrement; primaryKey; <-:create"`
	Synonym         string    `json:"synonym" validate:"required"`
	Class           string    `json:"class" validate:"required"`
	Subclass        string    `json:"subclass" validate:"required"`
	Storage         string    `json:"storage" validate:"required"`
	Incompatibility string    `json:"incompatibility" validate:"required"`
	Precautions     string    `json:"precautions" validate:"required"`
	Symbols         string    `json:"symbols" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Batch           string    `json:"batch" validate:"required"`
	Due_date        string    `json:"due_date" validate:"required"`
	Location        string    `json:"location" validate:"required"`
	Quantity        string    `json:"quantity" validate:"required"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
}
