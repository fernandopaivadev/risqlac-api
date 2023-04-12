package application

import (
	"risqlac-api/application/controllers"

	"github.com/labstack/echo/v4"
)

func (server *server) LoadDefaultRoutes() {
	server.Instance.GET("/", func(context echo.Context) error {
		return context.String(200, "RisQLAC API v2.5.15")
	})
}

func (server *server) LoadSessionRoutes() {
	sessionRoutes := server.Instance.Group("/session")

	sessionRoutes.GET(
		"/login",
		controllers.Session.Login,
	)
	sessionRoutes.GET(
		"/list",
		controllers.Session.List,
		Middleware.ValidateSessionToken,
	)
	sessionRoutes.DELETE(
		"/logout",
		controllers.Session.Logout,
		Middleware.ValidateSessionToken,
	)
	sessionRoutes.DELETE(
		"/complete-logout",
		controllers.Session.CompleteLogout,
		Middleware.ValidateSessionToken,
	)
}

func (server *server) LoadUserRoutes() {
	userRoutes := server.Instance.Group("/user")

	// userRoutes.GET(
	// 	"/request-password-change",
	// 	controllers.User.RequestPasswordChange,
	// )
	// userRoutes.GET(
	// 	"/change-password",
	// 	controllers.User.ChangePassword,
	// )
	userRoutes.POST(
		"/create",
		controllers.User.Create,
		Middleware.ValidateSessionToken,
	)
	userRoutes.PUT(
		"/update",
		controllers.User.Update,
		Middleware.ValidateSessionToken,
	)
	userRoutes.GET(
		"/list",
		controllers.User.List,
		Middleware.ValidateSessionToken,
	)
	userRoutes.DELETE(
		"/delete",
		controllers.User.Delete,
		Middleware.ValidateSessionToken,
	)
}

func (server *server) LoadProductRoutes() {
	productRoutes := server.Instance.Group("/product")

	productRoutes.POST(
		"/create",
		controllers.Product.Create,
		Middleware.ValidateSessionToken,
		Middleware.VerifyAdmin,
	)
	productRoutes.PUT(
		"/update",
		controllers.Product.Update,
		Middleware.ValidateSessionToken,
		Middleware.VerifyAdmin,
	)
	productRoutes.GET(
		"/list",
		controllers.Product.List,
		Middleware.ValidateSessionToken,
	)
	productRoutes.DELETE(
		"/delete",
		controllers.Product.Delete,
		Middleware.ValidateSessionToken,
		Middleware.VerifyAdmin,
	)
}

func (server *server) LoadReportRoutes() {
	reportRoutes := server.Instance.Group("/report")

	reportRoutes.GET(
		"/products/pdf",
		controllers.Report.GetProductsReportPDF,
		Middleware.ValidateSessionToken,
	)
	reportRoutes.GET(
		"/products/csv",
		controllers.Report.GetProductsReportCSV,
		Middleware.ValidateSessionToken,
	)
	reportRoutes.GET(
		"/products/xlsx",
		controllers.Report.GetProductsReportXLSX,
		Middleware.ValidateSessionToken,
	)
}
