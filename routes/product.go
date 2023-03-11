package routes

import (
	"risqlac-api/controllers"
	"risqlac-api/middleware"
	"risqlac-api/server"
)

func LoadProductRoutes() {
	productRoutes := server.Instance.Group("/product")

	productRoutes.Post(
		"/create",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.CreateProduct,
	)
	productRoutes.Put(
		"/update",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.UpdateProduct,
	)
	productRoutes.Get(
		"/list",
		middleware.ValidateToken,
		controllers.ListProducts,
	)
	productRoutes.Delete(
		"/delete",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.DeleteProduct,
	)
	productRoutes.Get(
		"/report",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.ProductReport,
	)
}
