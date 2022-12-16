package middleware

import (
	"risqlac-api/services"

	"github.com/gofiber/fiber/v2"
)

func ValidateToken(context *fiber.Ctx) error {
	headers := context.GetReqHeaders()
	tokenString := headers["Authorization"]

	isValid, err := services.ValidateUserToken(tokenString)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
			Message: "Error parsing token",
			Error:   err.Error(),
		})
	} else if isValid {
		return context.Next()
	} else {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
			Message: "Invalid token",
		})
	}
}
