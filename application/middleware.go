package application

import (
	"fmt"
	"risqlac-api/application/services"

	"github.com/labstack/echo/v4"
)

type middleware struct{}

var Middleware middleware

func (*middleware) ValidateSessionToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		headers := context.Request().Header
		tokenString := headers["Authorization"][0]
		user, err := services.User.ValidateSessionToken(tokenString)

		if err != nil {
			return context.JSON(400, echo.Map{
				"message": "session token validation error",
				"Error":   err.Error(),
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
		isAdmin := headers["Isadmin"][0] == "1"

		if !isAdmin {
			return context.JSON(403, echo.Map{
				"message": "not allowed for not admin users",
			})
		}

		return next(context)
	}
}
