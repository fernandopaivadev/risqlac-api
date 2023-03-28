package services

import (
	"errors"
	"risqlac-api/environment"

	"github.com/gofiber/fiber/v2"
)

type telegramBotService struct{}

var TelegramBot telegramBotService

var telegramBotApiUrl = "https://api.telegram.org/bot" + environment.Variables.TelegramBotApiToken

func (*telegramBotService) SendMessage(chatID string, message string) error {
	type RequestBody struct {
		ChatID string `json:"chat_id"`
		Text   string `json:"text"`
	}

	requestBody := RequestBody{
		ChatID: chatID,
		Text:   message,
	}

	url := telegramBotApiUrl + "/sendMessage"

	agent := fiber.Post(url)

	agent.JSON(requestBody)

	err := agent.Parse()

	if err != nil {
		return err
	}

	statusCode, _, _ := agent.Bytes()

	if statusCode != 200 {
		return errors.New("error sending message")
	}

	return nil
}

func (*telegramBotService) SendImageFromURL(chatID string, imageURL string) error {
	type RequestBody struct {
		ChatID string `json:"chat_id"`
		Photo  string `json:"photo"`
	}

	requestBody := RequestBody{
		ChatID: chatID,
		Photo:  imageURL,
	}

	url := telegramBotApiUrl + "/sendPhoto"

	agent := fiber.Post(url)

	agent.JSON(requestBody)

	err := agent.Parse()

	if err != nil {
		return err
	}

	statusCode, _, _ := agent.Bytes()

	if statusCode != 200 {
		return errors.New("error sending image")
	}

	return nil
}
