package application

import (
	"risqlac-api/application/controllers"

	"github.com/labstack/echo/v4"
)

func (server *server) LoadMetricsRoutes() {
	server.Instance.GET("/", func(context echo.Context) error {
		return context.String(200, "RisQLAC API v2.4.22")
	})
}

func (server *server) LoadUserRoutes() {
	userRoutes := server.Instance.Group("/user")

	userRoutes.GET(
		"/login",
		controllers.User.Login,
	)
	userRoutes.GET(
		"/request-password-change",
		controllers.User.RequestPasswordChange,
	)
	userRoutes.GET(
		"/change-password",
		controllers.User.ChangePassword,
	)
	userRoutes.POST(
		"/create",
		controllers.User.Create,
		// Middleware.ValidateToken,
	)
	userRoutes.PUT(
		"/update",
		controllers.User.Update,
		Middleware.ValidateToken,
	)
	userRoutes.GET(
		"/list",
		controllers.User.List,
		Middleware.ValidateToken,
	)
	userRoutes.DELETE(
		"/delete",
		controllers.User.Delete,
		Middleware.ValidateToken,
	)
}

func (server *server) LoadProductRoutes() {
	productRoutes := server.Instance.Group("/product")

	productRoutes.POST(
		"/create",
		controllers.Product.Create,
		Middleware.ValidateToken,
		Middleware.VerifyAdmin,
	)
	productRoutes.PUT(
		"/update",
		controllers.Product.Update,
		Middleware.ValidateToken,
		Middleware.VerifyAdmin,
	)
	productRoutes.GET(
		"/list",
		controllers.Product.List,
		Middleware.ValidateToken,
	)
	productRoutes.DELETE(
		"/delete",
		controllers.Product.Delete,
		Middleware.ValidateToken,
		Middleware.VerifyAdmin,
	)
	productRoutes.GET(
		"/report/pdf",
		controllers.Product.GetReportPDF,
		Middleware.ValidateToken,
	)
	productRoutes.GET(
		"/report/csv",
		controllers.Product.GetReportCSV,
		Middleware.ValidateToken,
	)
	productRoutes.GET(
		"/report/xlsx",
		controllers.Product.GetReportXLSX,
		Middleware.ValidateToken,
	)
}
