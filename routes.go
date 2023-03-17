package main

import (
	"risqlac-api/controllers"
)

func (server *Server) LoadUserRoutes() {
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

func (server *Server) LoadProductRoutes() {
	productRoutes := server.App.Group("/product")

	productRoutes.Post(
		"/create",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.Product.Create,
	)
	productRoutes.Put(
		"/update",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.Product.Update,
	)
	productRoutes.Get(
		"/list",
		middleware.ValidateToken,
		controllers.Product.List,
	)
	productRoutes.Delete(
		"/delete",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.Product.Delete,
	)
	productRoutes.Get(
		"/report/pdf",
		middleware.ValidateToken,
		controllers.Product.GetReportPDF,
	)
	productRoutes.Get(
		"/report/csv",
		middleware.ValidateToken,
		controllers.Product.GetReportCSV,
	)
	productRoutes.Get(
		"/report/xlsx",
		middleware.ValidateToken,
		controllers.Product.GetReportXLSX,
	)
}
