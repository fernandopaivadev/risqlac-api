package application

import (
	"fmt"
	"risqlac-api/application/services"

	"github.com/labstack/echo/v4"
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

func (*middleware) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		headers := context.Request().Header
		tokenString := headers["Authorization"][0]
		claims, err := services.Utils.ParseToken(tokenString)

		if err != nil {
			return context.JSON(401, errorResponse{
				Message: "Error validating token",
			})
		}

		user, err := services.User.GetById(claims.UserId)

		if err != nil {
			return context.JSON(401, errorResponse{
				Message: "Error retrieving user",
				Error:   err.Error(),
			})
		}

		context.Request().Header.Add("UserId", fmt.Sprint(user.Id))
		context.Request().Header.Add("IsAdmin", fmt.Sprint(user.IsAdmin))

		return next(context)
	}
}

func (*middleware) VerifyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		headers := context.Request().Header
		isAdmin := headers["Isadmin"][0] == "true"

		if !isAdmin {
			return context.JSON(403, messageResponse{
				Message: "Not allowed for no admin users",
			})
		}

		return next(context)
	}
}
