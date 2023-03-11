package controllers

import (
	"risqlac-api/reports"
	"risqlac-api/services"
	"risqlac-api/types"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(context *fiber.Ctx) error {
	var product types.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.ValidateStruct(product)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.CreateProduct(product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error creating product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(types.SuccessResponse{
		Message: "Product created",
	})
}

func UpdateProduct(context *fiber.Ctx) error {
	var product types.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.ValidateStruct(product)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.UpdateProduct(product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error updating product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "Product updated",
	})
}

func ListProducts(context *fiber.Ctx) error {
	var query types.ByIdRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	var products []types.Product

	if query.Id != 0 {
		product, err := services.GetProduct(query.Id)

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Message: "Error retrieving product",
				Error:   err.Error(),
			})
		}

		products = append(products, product)
	} else {
		products, err = services.ListProducts()

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Message: "Error retrieving products",
				Error:   err.Error(),
			})
		}
	}

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.ListProductsResponse{
		Products: products,
	})
}

func DeleteProduct(context *fiber.Ctx) error {
	var query types.ByIdRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.DeleteProduct(query.Id)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error deleting product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "Product deleted",
	})
}

func ProductReport(context *fiber.Ctx) error {
	products, err := services.ListProducts()

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	file, err := reports.GetProductsReport("listagem de produtos", products)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Error generating pdf",
			Error:   err.Error(),
		})
	}

	context.Response().Header.Set("Content-Type", "application/pdf")
	return context.Send(file)
}
