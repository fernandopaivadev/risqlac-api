package controllers

import "risqlac-api/models"

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type UserAuthResponse struct {
	Token string `json:"token"`
}

type ListUsersResponse struct {
	Users []models.User `json:"users"`
}

type ListProductsResponse struct {
	Products []models.Product `json:"products"`
}

type QueryById struct {
	Id uint64
}

type UserAuthQuery struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
