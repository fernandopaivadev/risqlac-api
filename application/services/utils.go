package services

import (
	"encoding/json"
	"risqlac-api/environment"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type utilsService struct{}

var Utils utilsService

func (*utilsService) GenerateJWT(userId uint64, expiresAt int64) (string, error) {
	claims := jwt.MapClaims{
		"UserId":    userId,
		"ExpiresAt": expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(environment.Variables.JwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (*utilsService) ParseToken(tokenString string) (tokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(environment.Variables.JwtSecret), nil
	})

	if err != nil {
		return tokenClaims{}, err
	}

	if !token.Valid {
		return tokenClaims{}, errors.New("invalid token")
	}

	jwtClaims, _ := token.Claims.(jwt.MapClaims)

	var claimsObject tokenClaims

	claimsJSON, _ := json.Marshal(jwtClaims)

	err = json.Unmarshal(claimsJSON, &claimsObject)

	if err != nil {
		return tokenClaims{}, err
	}

	return claimsObject, nil
}

func (*utilsService) ValidateStruct(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}

func (*utilsService) SendEmail(
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
