package routes

import (
	"github.com/gofiber/fiber/v2"
	"risqlac-api/controllers"
	"risqlac-api/middleware"
)

type UserRoutes struct{}

var User UserRoutes

func (routes *UserRoutes) Load(app *fiber.App) {
	userRoutes := app.Group("/user")

	userRoutes.Get(
		"/login",
		controllers.User.Login,
	)
	userRoutes.Get(
		"/request-password-change",
		controllers.User.RequestPasswordChange,
	)
	userRoutes.Get(
		"/change-password",
		middleware.ValidateToken,
		controllers.User.ChangePassword,
	)
	userRoutes.Post(
		"/create",
		// middleware.ValidateToken,
		controllers.User.Create,
	)
	userRoutes.Put(
		"/update",
		middleware.ValidateToken,
		controllers.User.Update,
	)
	userRoutes.Get(
		"/list",
		middleware.ValidateToken,
		controllers.User.List,
	)
	userRoutes.Delete(
		"/delete",
		middleware.ValidateToken,
		controllers.User.Delete,
	)
}
