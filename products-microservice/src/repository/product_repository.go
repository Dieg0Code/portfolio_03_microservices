package repository

import "github.com/dieg0code/products-microservice/src/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductById(ProductID uint) (*models.Product, error)
	GetAllProducts(offset int, pageSize int) ([]models.Product, error)
	GetByCategory(category string) ([]models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(ProductID uint) error
	CheckProductExist(ProductID uint) (bool, error)
}
