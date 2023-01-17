package routes

import (
	"risqlac-api/controllers"
	"risqlac-api/middleware"
)

func User() {
	userRoutes := App.Group("/user")

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
		middleware.ValidateToken,
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
