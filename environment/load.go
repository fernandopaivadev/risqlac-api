package environment

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type EnvVars struct {
	ServerPort     string
	JwtSecret      string
	DatabaseFile   string
	SendgridApiKey string
}

var Variables EnvVars

func Load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(
			"Error loading environment variables from .Variables file => " + err.Error(),
		)
	}

	Variables.DatabaseFile = os.Getenv("DATABASE_FILE")
	Variables.JwtSecret = os.Getenv("JWT_SECRET")
	Variables.SendgridApiKey = os.Getenv("SENDGRID_API_KEY")
	Variables.ServerPort = os.Getenv("SERVER_PORT")
}
