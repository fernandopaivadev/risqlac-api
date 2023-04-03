package application

import (
	"fmt"
	"risqlac-api/application/services"

	"github.com/labstack/echo/v4"
)

type middleware struct{}

var Middleware middleware

func (*middleware) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		headers := context.Request().Header
		tokenString := headers["Authorization"][0]
		claims, err := services.Utils.ParseToken(tokenString)

		if err != nil {
			return context.JSON(400, echo.Map{
				"message": "token validation error",
				"Error":   err.Error(),
			})
		}

		user, err := services.User.GetById(claims.UserId)

		if err != nil {
			return context.JSON(500, echo.Map{
				"message": "error retrieving user",
				"error":   err.Error(),
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
			return context.JSON(403, echo.Map{
				"message": "not allowed for not admin users",
			})
		}

		return next(context)
	}
}
