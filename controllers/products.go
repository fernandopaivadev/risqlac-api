package controllers

import (
	"risqlac-api/database"
	"risqlac-api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(context *fiber.Ctx) error {
	var product models.Product
	err := context.BodyParser(&product)

	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err,
		})
		return err
	}

	return context.Status(fiber.StatusCreated).JSON(CreatedProductResponse{
		CreatedProduct: product,
	})
}

func ListProducts(context *fiber.Ctx) error {
	return context.Status(fiber.StatusCreated).JSON(ListProductsResponse{
		Products: database.Products,
	})
}
