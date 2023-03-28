package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envVars struct {
	ServerPort          string
	JwtSecret           string
	DatabaseUrl         string
	TelegramBotApiToken string
	SendgridApiKey      string
}

var Variables envVars

func Load() {
	err := godotenv.Load()

	if err == nil {
		log.Println(
			"environment variables loaded from .env file",
		)
	}

	Variables.ServerPort = os.Getenv("SERVER_PORT")
	Variables.JwtSecret = os.Getenv("JWT_SECRET")
	Variables.DatabaseUrl = os.Getenv("DATABASE_URL")
	Variables.TelegramBotApiToken = os.Getenv("TELEGRAM_API_TOKEN")
	Variables.SendgridApiKey = os.Getenv("SENDGRID_API_KEY")
}
