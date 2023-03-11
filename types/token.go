package types

type TokenClaims struct {
	UserId    uint64 `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
}
