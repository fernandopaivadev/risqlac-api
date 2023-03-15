package services

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"risqlac-api/types"
	"time"
)

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

func GetProductsReport(products []types.Product) ([]byte, error) {
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
