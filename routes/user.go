package routes

import (
	"risqlac-api/app"
	"risqlac-api/controllers"
	"risqlac-api/middleware"
)

func SetupUserRoutes() {
	userRoutes := app.Instance.Group("/user")

	userRoutes.Get(
		"/login",
		controllers.UserLogin,
	)
	userRoutes.Get(
		"/request-password-change",
		controllers.RequestPasswordChange,
	)
	userRoutes.Get(
		"/change-password",
		middleware.ValidateToken,
		controllers.ChangePassword,
	)
	userRoutes.Post(
		"/create",
		// middleware.ValidateToken,
		controllers.CreateUser,
	)
	userRoutes.Put(
		"/update",
		middleware.ValidateToken,
		controllers.UpdateUser,
	)
	userRoutes.Get(
		"/list",
		middleware.ValidateToken,
		controllers.ListUsers,
	)
	userRoutes.Delete(
		"/delete",
		middleware.ValidateToken,
		controllers.DeleteUser,
	)
}
