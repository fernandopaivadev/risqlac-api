package services

import (
	"risqlac-api/application/models"
	"risqlac-api/infrastructure"
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
