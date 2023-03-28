package controllers

import "risqlac-api/application/models"

type messageResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type userAuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type userAuthResponse struct {
	Token string `json:"token"`
}

type requestPasswordChangeRequest struct {
	Email string `json:"email" validate:"required"`
}

type changePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

type byIdRequest struct {
	Id uint64 `json:"id"`
}

type listUsersResponse struct {
	Users []models.User `json:"users"`
}

type listProductsResponse struct {
	Products []models.Product `json:"products"`
}
