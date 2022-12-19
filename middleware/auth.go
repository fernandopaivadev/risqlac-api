package middleware

import (
	"encoding/json"
	"fmt"
	"risqlac-api/services"

	"github.com/gofiber/fiber/v2"
)

func ValidateToken(context *fiber.Ctx) error {
	headers := context.GetReqHeaders()
	tokenString := headers["Authorization"]
	isValid, claims, err := services.ValidateUserToken(tokenString)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
			Message: "Error parsing token",
			Error:   err.Error(),
		})
	}

	if !isValid {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
			Message: "Invalid token",
		})
	}

	claimsObject := TokenClaims{}
	claimsJSON, _ := json.Marshal(claims)
	err = json.Unmarshal(claimsJSON, &claimsObject)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
			Message: "Error parsing claims",
			Error:   err.Error(),
		})
	}

	users, err := services.ListUsers(claimsObject.User_Id)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	if len(users) <= 0 {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
			Message: "User not found",
		})
	}

	context.Request().Header.Add("user_id", fmt.Sprint(users[0].Id))
	context.Request().Header.Add("is_admin", fmt.Sprint(users[0].Is_admin))
	return context.Next()
}
