package application

import (
	"fmt"
	"risqlac-api/application/services"

	"github.com/gofiber/fiber/v2"
)

type middleware struct{}

var Middleware middleware

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type messageResponse struct {
	Message string `json:"message"`
}

func (*middleware) ValidateToken(context *fiber.Ctx) error {
	headers := context.GetReqHeaders()
	tokenString := headers["Authorization"]
	claims, err := services.Utils.ParseToken(tokenString)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(errorResponse{
			Message: "Error validating token",
			Error:   err.Error(),
		})
	}

	user, err := services.User.GetById(claims.UserId)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(errorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Error retrieving user",
			"error":   err.Error(),
		})
	}

	context.Request().Header.Add("UserId", fmt.Sprint(user.Id))
	context.Request().Header.Add("IsAdmin", fmt.Sprint(user.IsAdmin))

	return context.Next()
}

func (*middleware) VerifyAdmin(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Isadmin"] == "true"

	if !isAdmin {
		return context.Status(fiber.StatusForbidden).JSON(messageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	return context.Next()
}
