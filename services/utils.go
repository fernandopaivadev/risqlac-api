package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"risqlac-api/environment"
)

type UtilsService struct{}

var Utils UtilsService

func (_ *UtilsService) ValidateStruct(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}

func (_ *UtilsService) SendEmail(
	receiverName string,
	receiverEmailAddress string,
	subject string,
	plainTextContent string,
	htmlContent string,
) error {
	from := mail.NewEmail("RisQLAC", "risqlac@protonmail.com")
	to := mail.NewEmail(receiverName, receiverEmailAddress)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(environment.Variables.SendgridApiKey)
	response, err := client.Send(message)

	if err != nil {
		return err
	}

	statusCode := response.StatusCode

	if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		return nil
	}

	return errors.New("email not sent")
}
