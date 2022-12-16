package services

import (
	"risqlac-api/database"
	"risqlac-api/models"

	"gorm.io/gorm"
)

func CreateProduct(product models.Product) error {
	result := database.Instance.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateProduct(user models.Product) error {
	result := database.Instance.Model(&user).Updates(models.Product{
		Name:            user.Name,
		Synonym:         user.Synonym,
		Class:           user.Class,
		Subclass:        user.Subclass,
		Storage:         user.Storage,
		Incompatibility: user.Incompatibility,
		Precautions:     user.Precautions,
		Symbols:         user.Symbols,
		Batch:           user.Batch,
		Due_date:        user.Due_date,
		Location:        user.Location,
		Quantity:        user.Quantity,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ListProducts(productId uint64) ([]models.Product, error) {
	var products []models.Product
	var result *gorm.DB

	if productId == 0 {
		result = database.Instance.Find(&products)
	} else {
		result = database.Instance.Find(&products, []uint64{productId})
	}

	if result.Error != nil {
		return []models.Product{}, result.Error
	}

	return products, nil
}

func DeleteProduct(productId uint64) error {
	result := database.Instance.Delete(&models.Product{}, productId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
