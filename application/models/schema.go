package models

import (
	"time"
)

type Session struct {
	Id        uint64    `json:"id" gorm:"unique; autoIncrement; primaryKey; <-:create"`
	Token     string    `json:"token" gorm:"unique" validate:"required"`
	UserId    uint64    `json:"user_id" validate:"required"`
	ExpiresAt time.Time `json:"expires_at" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Id        uint64    `json:"id" gorm:"unique; autoIncrement; primaryKey; <-:create"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`
	Name      string    `json:"name" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	IsAdmin   uint8     `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	Id              uint64    `json:"id" gorm:"autoIncrement; primaryKey; <-:create"`
	Synonym         string    `json:"synonym" validate:"required"`
	Class           string    `json:"class" validate:"required"`
	Subclass        string    `json:"subclass" validate:"required"`
	Storage         string    `json:"storage" validate:"required"`
	Incompatibility string    `json:"incompatibility" validate:"required"`
	Precautions     string    `json:"precautions" validate:"required"`
	Symbols         string    `json:"symbols"`
	Name            string    `json:"name" validate:"required"`
	Batch           string    `json:"batch" validate:"required"`
	DueDate         string    `json:"due_date" validate:"required"`
	Location        string    `json:"location" validate:"required"`
	Quantity        string    `json:"quantity" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
