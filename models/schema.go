package models

import (
	"time"
)

type User struct {
	Id         uint64    `json:"id" gorm:"unique; autoIncrement; primaryKey; <-:create"`
	Username   string    `json:"username" gorm:"unique"`
	Email      string    `json:"email" gorm:"unique"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	Is_admin   bool      `json:"is_admin"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Product struct {
	Id              uint64    `json:"id" gorm:"autoIncrement; primaryKey; <-:create"`
	Synonym         string    `json:"synonym"`
	Class           string    `json:"class"`
	Subclass        string    `json:"subclass"`
	Storage         string    `json:"storage"`
	Incompatibility string    `json:"incompatibility"`
	Precautions     string    `json:"precautions"`
	Symbols         string    `json:"symbols"`
	Name            string    `json:"name"`
	Batch           string    `json:"batch"`
	Due_date        string    `json:"due_date"`
	Location        string    `json:"location"`
	Quantity        string    `json:"quantity"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
}
