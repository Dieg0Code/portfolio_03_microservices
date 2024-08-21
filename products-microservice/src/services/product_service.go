package services

import (
	"github.com/dieg0code/products-microservice/src/json/request"
	"github.com/dieg0code/products-microservice/src/json/response"
)

type ProductService interface {
	CreateProduct(product *request.CreateProductRequest) (*uint, error)
	GetProductById(productID uint) (*response.ProductResponse, error)
	GetAllProducts(page int, pageSize int) ([]response.ProductResponse, error)
	GetByCategory(category string) ([]response.ProductResponse, error)
	UpdateProduct(product *request.UpdateProductRequest) (*response.ProductResponse, error)
	DeleteProduct(ProductID uint) error
}
