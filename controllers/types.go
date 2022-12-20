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

type UserAuthQuery struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ListUsersResponse struct {
	Users []models.User `json:"users"`
}

type ListProductsResponse struct {
	Products []models.Product `json:"products"`
}

type ListUsersQuery struct {
	UserId uint64 `json:"user_id"`
}

type ListProdcutsQuery struct {
	ProductId uint64 `json:"product_id"`
}

type DeleteUserQuery struct {
	UserId uint64 `json:"user_id"`
}

type DeleteProdcutQuery struct {
	ProductId uint64 `json:"product_id"`
}
