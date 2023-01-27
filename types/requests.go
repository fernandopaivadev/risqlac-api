package types

type UserAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ByIdRequest struct {
	Id uint64 `json:"id"`
}

type RequestPasswordChangeRequest struct {
	Email string `json:"email"`
}

type ChangePasswordRequest struct {
	Password string `json:"password"`
}
