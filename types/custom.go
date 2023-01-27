package types

type EnvironmentVariables struct {
	SERVER_PORT      string
	JWT_SECRET       string
	DATABASE_FILE    string
	SENDGRID_API_KEY string
}

type TokenClaims struct {
	User_id    uint64 `json:"user_id"`
	Expires_at int64  `json:"expires_at"`
}
