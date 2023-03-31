package controllers

import (
	"risqlac-api/application/models"
	"risqlac-api/application/services"

	"github.com/gofiber/fiber/v2"
)

type productController struct{}

var Product productController

func (*productController) Create(context *fiber.Ctx) error {
	var product models.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(product)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.Product.Create(product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error creating product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(messageResponse{
		Message: "Product created",
	})
}

func (*productController) Update(context *fiber.Ctx) error {
	var product models.Product
	err := context.BodyParser(&product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(product)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.Product.Update(product)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error updating product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(messageResponse{
		Message: "Product updated",
	})
}

func (*productController) List(context *fiber.Ctx) error {
	var query byIdRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	var products []models.Product

	if query.Id != 0 {
		product, err := services.Product.GetById(query.Id)

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
				Message: "Error retrieving product",
				Error:   err.Error(),
			})
		}

		products = append(products, product)
	} else {
		products, err = services.Product.List()

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
				Message: "Error retrieving products",
				Error:   err.Error(),
			})
		}
	}

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(listProductsResponse{
		Products: products,
	})
}

func (*productController) Delete(context *fiber.Ctx) error {
	var query byIdRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.Product.Delete(query.Id)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error deleting product",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(messageResponse{
		Message: "Product deleted",
	})
}

func (*productController) GetReportPDF(context *fiber.Ctx) error {
	products, err := services.Product.List()

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	file, err := services.Product.GetReportPDF(products)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Error generating pdf",
			Error:   err.Error(),
		})
	}

	context.Response().Header.Set("Content-Type", "application/pdf")
	return context.Send(file)
}

func (*productController) GetReportCSV(context *fiber.Ctx) error {
	products, err := services.Product.List()

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	file, err := services.Product.GetReportCSV(products)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Error generating csv",
			Error:   err.Error(),
		})
	}

	context.Response().Header.Set("Content-Type", "application/csv")
	return context.Send(file)
}

func (*productController) GetReportXLSX(context *fiber.Ctx) error {
	products, err := services.Product.List()

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}

	file, err := services.Product.GetReportXLSX(products)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Error generating xlsx",
			Error:   err.Error(),
		})
	}

	context.Response().Header.Set(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	)
	return context.Send(file)
}