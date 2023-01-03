package environment

import (
	"os"
	"risqlac-api/types"

	"github.com/joho/godotenv"
)

var env types.Env

func Load() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading environment variables")
	}

	env.DATABASE_FILE = os.Getenv("DATABASE_FILE")
	env.JWT_SECRET = os.Getenv("JWT_SECRET")
	env.SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")
}

func Get() types.Env {
	return env
}
