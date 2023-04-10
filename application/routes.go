package application

import (
	"risqlac-api/application/controllers"

	"github.com/labstack/echo/v4"
)

func (server *server) LoadDefaultRoutes() {
	server.Instance.GET("/", func(context echo.Context) error {
		return context.String(200, "RisQLAC API v2.5.12")
	})
}

func (server *server) LoadUserRoutes() {
	userRoutes := server.Instance.Group("/user")

	userRoutes.GET(
		"/login",
		controllers.User.Login,
	)
	userRoutes.GET(
		"/sessions",
		controllers.User.ListSessions,
		Middleware.ValidateSessionToken,
	)
	userRoutes.DELETE(
		"/logout",
		controllers.User.Logout,
		Middleware.ValidateSessionToken,
	)
	userRoutes.DELETE(
		"/complete-logout",
		controllers.User.CompleteLogout,
		Middleware.ValidateSessionToken,
	)
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
	productRoutes.GET(
		"/report/pdf",
		controllers.Product.GetReportPDF,
		Middleware.ValidateSessionToken,
	)
	productRoutes.GET(
		"/report/csv",
		controllers.Product.GetReportCSV,
		Middleware.ValidateSessionToken,
	)
	productRoutes.GET(
		"/report/xlsx",
		controllers.Product.GetReportXLSX,
		Middleware.ValidateSessionToken,
	)
}
