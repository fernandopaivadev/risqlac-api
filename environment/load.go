package environment

import (
	"fmt"
	"os"
	"risqlac-api/types"

	"github.com/joho/godotenv"
)

var env types.EnvironmentVariables

func Load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(
			"Error loading environment variables from .env file => " + err.Error(),
		)
	}

	env.DATABASE_FILE = os.Getenv("DATABASE_FILE")
	env.JWT_SECRET = os.Getenv("JWT_SECRET")
	env.SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")
	env.SERVER_PORT = os.Getenv("SERVER_PORT")
}
