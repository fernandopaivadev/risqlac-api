package controllers

import (
	"risqlac-api/application/services"

	"github.com/labstack/echo/v4"
)

type reportController struct{}

var Report reportController

func (*reportController) GetProductsReportPDF(context echo.Context) error {
	products, err := services.Product.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving products",
			"error":   err.Error(),
		})
	}

	file, err := services.Report.GetProductsReportPDF(products)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error generating pdf",
			"error":   err.Error(),
		})
	}

	return context.Blob(200, "application/pdf", file)
}

func (*reportController) GetProductsReportCSV(context echo.Context) error {
	products, err := services.Product.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving products",
			"error":   err.Error(),
		})
	}

	file, err := services.Report.GetProductsReportCSV(products)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error generating csv",
			"error":   err.Error(),
		})
	}

	return context.Blob(200, "text/csv", file)
}

func (*reportController) GetProductsReportXLSX(context echo.Context) error {
	products, err := services.Product.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving products",
			"error":   err.Error(),
		})
	}

	file, err := services.Report.GetProductsReportXLSX(products)

	if err != nil {
		return context.JSON(500, echo.Map{
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
