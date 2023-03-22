package main

import (
	"fmt"
	"risqlac-api/services"
	"risqlac-api/types"

	"github.com/gofiber/fiber/v2"
)

type middleware struct{}

var Middleware middleware

func (*middleware) ValidateToken(context *fiber.Ctx) error {
	headers := context.GetReqHeaders()
	tokenString := headers["Authorization"]
	claims, err := services.Utils.ParseToken(tokenString)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(types.ErrorResponse{
			Message: "Error validating token",
			Error:   err.Error(),
		})
	}

	user, err := services.User.GetById(claims.UserId)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(types.ErrorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	context.Request().Header.Add("UserId", fmt.Sprint(user.Id))
	context.Request().Header.Add("IsAdmin", fmt.Sprint(user.IsAdmin))

	return context.Next()
}

func (*middleware) VerifyAdmin(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["IsAdmin"] == "true"

	if !isAdmin {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	return context.Next()
}
