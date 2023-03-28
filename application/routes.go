package application

import (
	"risqlac-api/application/controllers"
)

func (server *server) LoadUserRoutes() {
	userRoutes := server.App.Group("/user")

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
		controllers.User.ChangePassword,
	)
	userRoutes.Post(
		"/create",
		// Middleware.ValidateToken,
		controllers.User.Create,
	)
	userRoutes.Put(
		"/update",
		Middleware.ValidateToken,
		controllers.User.Update,
	)
	userRoutes.Get(
		"/list",
		Middleware.ValidateToken,
		controllers.User.List,
	)
	userRoutes.Delete(
		"/delete",
		Middleware.ValidateToken,
		controllers.User.Delete,
	)
}

func (server *server) LoadProductRoutes() {
	productRoutes := server.App.Group("/product")

	productRoutes.Post(
		"/create",
		Middleware.ValidateToken,
		Middleware.VerifyAdmin,
		controllers.Product.Create,
	)
	productRoutes.Put(
		"/update",
		Middleware.ValidateToken,
		Middleware.VerifyAdmin,
		controllers.Product.Update,
	)
	productRoutes.Get(
		"/list",
		Middleware.ValidateToken,
		controllers.Product.List,
	)
	productRoutes.Delete(
		"/delete",
		Middleware.ValidateToken,
		Middleware.VerifyAdmin,
		controllers.Product.Delete,
	)
	productRoutes.Get(
		"/report/pdf",
		Middleware.ValidateToken,
		controllers.Product.GetReportPDF,
	)
	productRoutes.Get(
		"/report/csv",
		Middleware.ValidateToken,
		controllers.Product.GetReportCSV,
	)
	productRoutes.Get(
		"/report/xlsx",
		Middleware.ValidateToken,
		controllers.Product.GetReportXLSX,
	)
}
