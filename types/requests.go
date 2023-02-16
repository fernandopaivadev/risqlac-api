package types

type UserAuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ByIdRequest struct {
	Id uint64 `json:"id"`
}

type RequestPasswordChangeRequest struct {
	Email string `json:"email" validate:"required"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}
