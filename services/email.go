package services

import (
	"errors"
	"risqlac-api/environment"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(
	receiverName string,
	receiverEmailAddress string,
	subject string,
	plainTextContent string,
	htmlContent string,
) error {
	from := mail.NewEmail("RisQLAC", "risqlac@protonmail.com")
	to := mail.NewEmail(receiverName, receiverEmailAddress)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(environment.Get().SENDGRID_API_KEY)
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
