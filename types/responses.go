package types

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type UserAuthResponse struct {
	Token string `json:"token"`
}

type ListUsersResponse struct {
	Users []User `json:"users"`
}

type ListProductsResponse struct {
	Products []Product `json:"products"`
}
