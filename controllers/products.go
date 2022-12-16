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
			Message: "Error parsing body params",
			Error:   err,
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

func DeleteProduct(context *fiber.Ctx) error {
	var query DeleteQuery
	err := context.QueryParser(&query)

	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "Error parsing query params",
			Error:   err,
		})
		return err
	}

	slice := database.Products
	var indexToDelete uint64

	for index, user := range slice {
		if user.Id == query.Id {
			indexToDelete = uint64(index)
		}
	}

	copy(slice[indexToDelete:], slice[indexToDelete+1:])
	slice = slice[:len(slice)-1]

	database.Products = slice

	return context.SendStatus(fiber.StatusOK)
}
