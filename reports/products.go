package reports

import (
	"risqlac-api/types"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func formatProductList(products []types.Product) ([]string, [][]string) {
	headers := []string{
		"Nome",
		"Sinônimos",
		"Classe",
		"Subclasse",
		"Armazenagem",
		"Incompat.",
		"Precauções",
		"Lote",
		"Validade",
		"Local",
		"Qtde.",
		"Data de cadastro",
	}
	var content [][]string

	for _, product := range products {
		row := []string{
			product.Name,
			product.Synonym,
			product.Class,
			product.Subclass,
			product.Storage,
			product.Incompatibility,
			product.Precautions,
			product.Batch,
			product.Due_date,
			product.Location,
			product.Quantity,
			product.Created_at.Local().String(),
		}
		content = append(content, row)
	}

	return headers, content
}

func GetProductsReport(fileName string, products []types.Product) ([]byte, error) {
	maroto := pdf.NewMaroto(consts.Landscape, consts.A4)
	maroto.SetPageMargins(1, 5, 5)

	headers, content := formatProductList(products)

	maroto.TableList(headers, content, props.TableList{
		HeaderProp: props.TableListContent{
			Family: consts.Arial,
			Style:  consts.Bold,
			Size:   8.0,
		},
		ContentProp: props.TableListContent{
			Family: consts.Courier,
			Style:  consts.Normal,
			Size:   6.0,
		},
		Align:              consts.Center,
		HeaderContentSpace: 1.0,
	})

	file, err := maroto.Output()

	if err != nil {
		return nil, err
	}

	return file.Bytes(), nil
}
