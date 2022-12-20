package middleware

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
