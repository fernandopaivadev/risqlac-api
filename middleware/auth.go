package middleware

import (
	"encoding/json"
	"fmt"
	"risqlac-api/services"
	"risqlac-api/types"

	"github.com/gofiber/fiber/v2"
)

func ValidateToken(context *fiber.Ctx) error {
	headers := context.GetReqHeaders()
	tokenString := headers["Authorization"]
	isValid, claims, err := services.ValidateUserToken(tokenString)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(types.ErrorResponse{
			Message: "Error parsing token",
			Error:   err.Error(),
		})
	}

	if !isValid {
		return context.Status(fiber.StatusUnauthorized).JSON(types.MessageResponse{
			Message: "Invalid token",
		})
	}

	var claimsObject types.TokenClaims
	claimsJSON, _ := json.Marshal(claims)
	err = json.Unmarshal(claimsJSON, &claimsObject)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(types.ErrorResponse{
			Message: "Error parsing claims",
			Error:   err.Error(),
		})
	}

	user, err := services.GetUser(claimsObject.User_Id)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(types.ErrorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	context.Request().Header.Add("user_id", fmt.Sprint(user.Id))
	context.Request().Header.Add("is_admin", fmt.Sprint(user.Is_admin))

	return context.Next()
}
