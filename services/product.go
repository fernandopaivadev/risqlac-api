package services

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"risqlac-api/database"
	"risqlac-api/types"
	"time"
)

type ProductService struct{}

var Product ProductService

func (service *ProductService) Create(product types.Product) error {
	result := database.Instance.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *ProductService) Update(user types.Product) error {
	result := database.Instance.Model(&user).Select("*").Updates(types.Product{
		Name:            user.Name,
		Synonym:         user.Synonym,
		Class:           user.Class,
		Subclass:        user.Subclass,
		Storage:         user.Storage,
		Incompatibility: user.Incompatibility,
		Precautions:     user.Precautions,
		Symbols:         user.Symbols,
		Batch:           user.Batch,
		DueDate:         user.DueDate,
		Location:        user.Location,
		Quantity:        user.Quantity,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *ProductService) GetById(productId uint64) (types.Product, error) {
	var product types.Product

	result := database.Instance.First(&product, productId)

	if result.Error != nil {
		return types.Product{}, result.Error
	}

	return product, nil
}

func (service *ProductService) List() ([]types.Product, error) {
	var products []types.Product

	result := database.Instance.Find(&products)

	if result.Error != nil {
		return []types.Product{}, result.Error
	}

	return products, nil
}

func (service *ProductService) Delete(productId uint64) error {
	result := database.Instance.Delete(&types.Product{}, productId)

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

func makeImage(maroto pdf.Maroto, path string) {
	maroto.Row(20.0, func() {
		maroto.Col(13.0, func() {
			_ = maroto.FileImage(path, props.Rect{
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

func (service *ProductService) GetReport(products []types.Product) ([]byte, error) {
	maroto := pdf.NewMaroto(consts.Portrait, consts.A4)
	maroto.SetPageMargins(20, 5, 20)
	maroto.SetTitle("Relatório de Produtos", true)

	makeImage(maroto, "./assets/logo.png")
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
