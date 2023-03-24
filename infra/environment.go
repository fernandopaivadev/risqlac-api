package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envVars struct {
	ServerPort          string
	JwtSecret           string
	DatabaseFile        string
	TelegramBotApiToken string
	SendgridApiKey      string
}

type environment struct {
	Variables envVars
}

var Environment environment

func (environment *environment) Load() {
	err := godotenv.Load()

	if err == nil {
		log.Println(
			"Environment variables loaded from .env file",
		)
	}

	environment.Variables.ServerPort = os.Getenv("SERVER_PORT")
	environment.Variables.JwtSecret = os.Getenv("JWT_SECRET")
	environment.Variables.DatabaseFile = os.Getenv("DATABASE_FILE")
	environment.Variables.TelegramBotApiToken = os.Getenv("TELEGRAM_API_TOKEN")
	environment.Variables.SendgridApiKey = os.Getenv("SENDGRID_API_KEY")
}
