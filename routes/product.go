package routes

import (
	"github.com/gofiber/fiber/v2"
	"risqlac-api/controllers"
	"risqlac-api/middleware"
)

type ProductRoutes struct{}

var Product ProductRoutes

func (routes *ProductRoutes) Load(app *fiber.App) {
	productRoutes := app.Group("/product")

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
		"/report",
		middleware.ValidateToken,
		controllers.Product.GetReport,
	)
}
