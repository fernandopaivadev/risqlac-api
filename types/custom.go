package types

import "risqlac-api/models"

type customError struct {
	Message string
}

func (err *customError) Error() string {
	return err.Message
}

func MakeCustomError(message string) error {
	return &customError{
		Message: message,
	}
}

type EnvironmentVariables struct {
	JWT_SECRET       string
	DATABASE_FILE    string
	SENDGRID_API_KEY string
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenClaims struct {
	User_id    uint64 `json:"user_id"`
	Expires_at int64  `json:"expires_at"`
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

type QueryById struct {
	Id uint64 `json:"id"`
}

type RequestPasswordChangeQuery struct {
	Email string `json:"email"`
}

type ChangePasswordQuery struct {
	Password string `json:"password"`
}
