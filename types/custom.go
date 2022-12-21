package types

import "risqlac-api/models"

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenClaims struct {
	User_Id    uint64 `json:"user_id"`
	Expires_at uint64 `json:"expires_at"`
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

type ListProductsQuery struct {
	ProductId uint64 `json:"product_id"`
}

type DeleteUserQuery struct {
	UserId uint64 `json:"user_id"`
}

type DeleteProductsQuery struct {
	ProductId uint64 `json:"product_id"`
}
