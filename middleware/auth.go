package middleware

import (
	"fmt"
	"risqlac-api/services"
	"risqlac-api/types"

	"github.com/gofiber/fiber/v2"
)

func ValidateToken(context *fiber.Ctx) error {
	headers := context.GetReqHeaders()
	tokenString := headers["Authorization"]
	claims, err := services.ParseUserToken(tokenString)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(types.ErrorResponse{
			Message: "Error validating token",
			Error:   err.Error(),
		})
	}

	user, err := services.GetUserById(claims.UserId)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(types.ErrorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	context.Request().Header.Add("user_id", fmt.Sprint(user.Id))
	context.Request().Header.Add("is_admin", fmt.Sprint(user.IsAdmin))

	return context.Next()
}
