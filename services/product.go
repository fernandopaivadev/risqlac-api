package services

import (
	"risqlac-api/database"
	"risqlac-api/types"
)

func CreateProduct(product types.Product) error {
	result := database.Instance.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateProduct(user types.Product) error {
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

func GetProduct(productId uint64) (types.Product, error) {
	var product types.Product

	result := database.Instance.First(&product, productId)

	if result.Error != nil {
		return types.Product{}, result.Error
	}

	return product, nil
}

func ListProducts() ([]types.Product, error) {
	var products []types.Product

	result := database.Instance.Find(&products)

	if result.Error != nil {
		return []types.Product{}, result.Error
	}

	return products, nil
}

func DeleteProduct(productId uint64) error {
	result := database.Instance.Delete(&types.Product{}, productId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
