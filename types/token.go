package types

type TokenClaims struct {
	UserId    uint64 `json:"UserId"`
	ExpiresAt int64  `json:"ExpiresAt"`
}
