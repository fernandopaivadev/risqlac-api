package controllers

import "risqlac-api/models"

type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type CreatedUserResponse struct {
	CreatedUser models.User `json:"created_user"`
}

type ListUsersResponse struct {
	Users []models.User `json:"users"`
}

type CreatedProductResponse struct {
	CreatedProduct models.Product `json:"created_product"`
}

type ListProductsResponse struct {
	Products []models.Product `json:"products"`
}

type DeleteQuery struct {
	Id uint64
}
