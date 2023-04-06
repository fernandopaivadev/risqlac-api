package services

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"log"
	"risqlac-api/application/assets"
	"risqlac-api/application/models"
	"risqlac-api/infrastructure"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/xuri/excelize/v2"
)

type productService struct{}

var Product productService

func (*productService) Create(product models.Product) error {
	result := infrastructure.Database.Instance.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*productService) Update(product models.Product) error {
	result := infrastructure.Database.Instance.Model(&product).Select("*").Updates(models.Product{
		Name:            product.Name,
		Synonym:         product.Synonym,
		Class:           product.Class,
		Subclass:        product.Subclass,
		Storage:         product.Storage,
		Incompatibility: product.Incompatibility,
		Precautions:     product.Precautions,
		Symbols:         product.Symbols,
		Batch:           product.Batch,
		DueDate:         product.DueDate,
		Location:        product.Location,
		Quantity:        product.Quantity,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*productService) GetById(productId uint64) (models.Product, error) {
	var product models.Product

	result := infrastructure.Database.Instance.First(&product, productId)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

func (*productService) List() ([]models.Product, error) {
	var products []models.Product

	result := infrastructure.Database.Instance.Find(&products)

	if result.Error != nil {
		return []models.Product{}, result.Error
	}

	return products, nil
}

func (*productService) Delete(productId uint64) error {
	result := infrastructure.Database.Instance.Delete(&models.Product{}, productId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func makeItem(maroto pdf.Maroto, title string, content string) {
	maroto.Row(5, func() {
		maroto.Col(3, func() {
			maroto.Text(title, props.Text{
				Family: consts.Arial,
				Style:  consts.Bold,
				Size:   10,
			})
		})
		maroto.Col(9, func() {
			maroto.Text(content, props.Text{
				Family: consts.Courier,
				Style:  consts.Normal,
				Size:   10,
			})
		})
	})
}

func makeTitle(maroto pdf.Maroto, title string) {
	maroto.Row(10, func() {
		maroto.Col(13, func() {
			maroto.Text(title, props.Text{
				Family: consts.Arial,
				Style:  consts.Bold,
				Size:   16,
				Align:  consts.Center,
			})
		})
	})
}

func makeImage(maroto pdf.Maroto, image []byte) {
	base64Image := base64.StdEncoding.EncodeToString(image)

	maroto.Row(20.0, func() {
		maroto.Col(13.0, func() {
			_ = maroto.Base64Image(base64Image, consts.Png, props.Rect{
				Left:    0,
				Top:     0,
				Percent: 0,
				Center:  true,
			})
		})
	})
}

func makeSpacer(maroto pdf.Maroto) {
	maroto.Line(5)
}

func formatDatetime(datetime time.Time) string {
	return datetime.Local().Format("02/01/2006")
}

func (*productService) GetReportPDF(products []models.Product) ([]byte, error) {
	maroto := pdf.NewMaroto(consts.Portrait, consts.A4)
	maroto.SetPageMargins(20, 5, 20)
	maroto.SetTitle("Relatório de Produtos", true)

	makeImage(maroto, assets.LogoRisQLAC)
	makeTitle(maroto, "Relatório de Produtos")

	for i := range products {
		makeSpacer(maroto)
		makeItem(maroto, "Nome:", products[i].Name)
		makeItem(maroto, "Sinônimo:", products[i].Synonym)
		makeItem(maroto, "Classe:", products[i].Class)
		makeItem(maroto, "Subclasse:", products[i].Subclass)
		makeItem(maroto, "Armazenagem:", products[i].Storage)
		makeItem(maroto, "Imcompatibilidade:", products[i].Incompatibility)
		makeItem(maroto, "Precauções:", products[i].Precautions)
		makeItem(maroto, "Lote:", products[i].Batch)
		makeItem(maroto, "Local:", products[i].Location)
		makeItem(maroto, "Quantidade:", products[i].Quantity)
		makeItem(maroto, "Data de cadastro:", formatDatetime(products[i].CreatedAt))
	}

	file, err := maroto.Output()

	if err != nil {
		return nil, err
	}

	return file.Bytes(), nil
}

func (*productService) GetReportCSV(products []models.Product) ([]byte, error) {
	rows := [][]string{
		{
			"Nome",
			"Sinônimo",
			"Classe",
			"Subclasse",
			"Armazenagem",
			"Imcompatibilidade",
			"Precauções",
			"Lote",
			"Local",
			"Quantidade",
			"Data de cadastro",
		},
	}

	for i := range products {
		rows = append(rows, []string{
			products[i].Name,
			products[i].Synonym,
			products[i].Class,
			products[i].Subclass,
			products[i].Storage,
			products[i].Incompatibility,
			products[i].Precautions,
			products[i].Batch,
			products[i].Location,
			products[i].Quantity,
			formatDatetime(products[i].CreatedAt),
		})
	}

	var buffer bytes.Buffer
	err := csv.NewWriter(&buffer).WriteAll(rows)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (*productService) GetReportXLSX(products []models.Product) ([]byte, error) {
	rows := [][]string{
		{
			"Nome",
			"Sinônimos",
			"Classe",
			"Subclasse",
			"Armazenagem",
			"Incompatibilidades",
			"Precauções",
			"Lote",
			"Local",
			"Quantidade",
			"Data de cadastro",
		},
	}

	for i := range products {
		rows = append(rows, []string{
			products[i].Name,
			products[i].Synonym,
			products[i].Class,
			products[i].Subclass,
			products[i].Storage,
			products[i].Incompatibility,
			products[i].Precautions,
			products[i].Batch,
			products[i].Location,
			products[i].Quantity,
			formatDatetime(products[i].CreatedAt),
		})
	}

	file := excelize.NewFile()

	defer func() {
		err := file.Close()

		if err != nil {
			log.Println(err)
		}
	}()

	file.SetActiveSheet(0)

	sheetName := "Lista de produtos"

	err := file.SetSheetName("Sheet1", sheetName)

	if err != nil {
		return nil, err
	}

	for i := range rows {
		cell := "A" + strconv.Itoa(i+1)

		err := file.SetSheetRow(sheetName, cell, &rows[i])

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	buffer, err := file.WriteToBuffer()

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
