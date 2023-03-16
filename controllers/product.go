package controllers

import (
	"risqlac-api/services"
	"risqlac-api/types"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct{}

var Product ProductController

func (controller *ProductController) Create(context *fiber.Ctx) error {
	var product types.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(product)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.Product.Create(product)

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

func (controller *ProductController) Update(context *fiber.Ctx) error {
	var product types.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(product)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.Product.Update(product)

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

func (controller *ProductController) List(context *fiber.Ctx) error {
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
		product, err := services.Product.GetById(query.Id)

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Message: "Error retrieving product",
				Error:   err.Error(),
			})
		}

		products = append(products, product)
	} else {
		products, err = services.Product.List()

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

func (controller *ProductController) Delete(context *fiber.Ctx) error {
	var query types.ByIdRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.Product.Delete(query.Id)

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

func (controller *ProductController) GetReport(context *fiber.Ctx) error {
	products, err := services.Product.List()

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	file, err := services.Product.GetReport(products)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Error generating pdf",
			Error:   err.Error(),
		})
	}

	context.Response().Header.Set("Content-Type", "application/pdf")
	return context.Send(file)
}
