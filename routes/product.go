package routes

import (
	"risqlac-api/controllers"
	"risqlac-api/middleware"
)

func Product() {
	productRoutes := App.Group("/product")

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
}
