package controllers

import (
	"risqlac-api/models"
	"risqlac-api/services"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(context *fiber.Ctx) error {
	var product models.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.CreateProduct(product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error creating product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Message: "Product created",
	})
}

func UpdateProduct(context *fiber.Ctx) error {
	var product models.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.UpdateProduct(product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error updating product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Product updated",
	})
}

func ListProducts(context *fiber.Ctx) error {
	var query ListProdcutsQuery
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	product, err := services.GetProduct(query.ProductId)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(ListProductsResponse{
		Products: []models.Product{product},
	})
}

func DeleteProduct(context *fiber.Ctx) error {
	var query DeleteProdcutQuery
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.DeleteProduct(query.ProductId)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error deleting product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Product deleted",
	})
}
