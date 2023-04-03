package controllers

import (
	"risqlac-api/application/models"
	"risqlac-api/application/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productController struct{}

var Product productController

func (*productController) Create(context echo.Context) error {
	var product models.Product
	err := context.Bind(&product)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing body",
			"error":   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(product)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	err = services.Product.Create(product)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error creating product",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "product created",
	})
}

func (*productController) Update(context echo.Context) error {
	var product models.Product
	err := context.Bind(&product)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing body",
			"error":   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(product)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	err = services.Product.Update(product)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error updating product",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "product updated",
	})
}

func (*productController) List(context echo.Context) error {
	productId, _ := strconv.ParseUint(context.QueryParam("id"), 10, 64)

	if productId != 0 {
		product, err := services.Product.GetById(productId)

		if err != nil {
			return context.JSON(500, echo.Map{
				"message": "error retrieving product",
				"error":   err.Error(),
			})
		}

		return context.JSON(200, echo.Map{
			"products": []models.Product{product},
		})
	}

	products, err := services.Product.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving products",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"products": products,
	})
}

func (*productController) Delete(context echo.Context) error {
	productId, err := strconv.ParseUint(context.QueryParam("id"), 10, 64)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   "id must be a number",
		})
	}

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	err = services.Product.Delete(productId)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error deleting product",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "product deleted",
	})
}

func (*productController) GetReportPDF(context echo.Context) error {
	products, err := services.Product.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving products",
			"error":   err.Error(),
		})
	}

	file, err := services.Product.GetReportPDF(products)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "error generating pdf",
			"error":   err.Error(),
		})
	}

	return context.Blob(200, "application/pdf", file)
}

func (*productController) GetReportCSV(context echo.Context) error {
	products, err := services.Product.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving products",
			"error":   err.Error(),
		})
	}

	file, err := services.Product.GetReportCSV(products)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "error generating csv",
			"error":   err.Error(),
		})
	}

	return context.Blob(200, "text/csv", file)
}

func (*productController) GetReportXLSX(context echo.Context) error {
	products, err := services.Product.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving products",
			"error":   err.Error(),
		})
	}

	file, err := services.Product.GetReportXLSX(products)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "error generating xlsx",
			"error":   err.Error(),
		})
	}

	return context.Blob(
		200,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		file,
	)
}
