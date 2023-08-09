package models

import (
	"GORM_API/configs"
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Qty         int       `json:"qty"`
	LastUpdated time.Time `json:"last_updated"`
	CreatedAt   time.Time `json:"created_at"`
}

func CreateProduct(product *Product) {
	configs.DB.Create(product)
}

func GetProducts() ([]Product, error) {
	var products []Product
	result := configs.DB.Find(&products)
	return products, result.Error
}

func GetProductByCode(code string) (Product, error) {
	var product Product
	result := configs.DB.Where("code = ?", code).First(&product)
	return product, result.Error
}

func DeleteProductByID(id uint) error {
	result := configs.DB.Delete(&Product{}, id)
	return result.Error
}
