package services

import (
	"errors"

	"github.com/dieg0code/products-microservice/src/json/request"
	"github.com/dieg0code/products-microservice/src/json/response"
	"github.com/dieg0code/products-microservice/src/models"
	"github.com/dieg0code/products-microservice/src/repository"
	"github.com/sirupsen/logrus"
)

type ProductServiceImpl struct {
	productRepo repository.ProductRepository
}

// CreateProduct implements ProductService.
func (p *ProductServiceImpl) CreateProduct(product *request.CreateProductRequest) (*uint, error) {

	productModel := &models.Product{
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
		Stock:    product.Stock,
	}

	createdProduct, err := p.productRepo.CreateProduct(productModel)
	if err != nil {
		logrus.WithError(err).Error("Error creating product")
		return nil, err
	}

	return &createdProduct.ID, nil
}

// DeleteProduct implements ProductService.
func (p *ProductServiceImpl) DeleteProduct(productID uint) error {

	if productID == 0 {
		return errors.New("product id is required")
	}

	err := p.productRepo.DeleteProduct(productID)
	if err != nil {
		logrus.WithError(err).Error("Error deleting product")
		return err
	}

	return nil
}

// GetAllProducts implements ProductService.
func (p *ProductServiceImpl) GetAllProducts(page int, pageSize int) ([]response.ProductResponse, error) {

	offset := (page - 1) * pageSize

	products, err := p.productRepo.GetAllProducts(offset, pageSize)
	if err != nil {
		logrus.WithError(err).Error("Error getting all products")
		return nil, err
	}

	var productResponses []response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, response.ProductResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Category:  product.Category,
			Price:     product.Price,
			Stock:     product.Stock,
		})
	}

	return productResponses, nil
}

// GetByCategory implements ProductService.
func (p *ProductServiceImpl) GetByCategory(category string) ([]response.ProductResponse, error) {

	products, err := p.productRepo.GetByCategory(category)
	if err != nil {
		logrus.WithError(err).Error("Error getting products by category")
		return nil, err
	}

	var productResponses []response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, response.ProductResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Category:  product.Category,
			Price:     product.Price,
			Stock:     product.Stock,
		})
	}

	return productResponses, nil
}

// GetProductById implements ProductService.
func (p *ProductServiceImpl) GetProductById(ProductID uint) (*response.ProductResponse, error) {

	product, err := p.productRepo.GetProductById(ProductID)
	if err != nil {
		logrus.WithError(err).Error("Error getting product by ID")
		return nil, err
	}

	productResponse := &response.ProductResponse{
		ProductID: product.ID,
		Name:      product.Name,
		Category:  product.Category,
		Price:     product.Price,
		Stock:     product.Stock,
	}

	return productResponse, nil
}

// UpdateProduct implements ProductService.
func (p *ProductServiceImpl) UpdateProduct(product *request.UpdateProductRequest) (*response.ProductResponse, error) {

	productModel := &models.Product{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}

	updatedProduct, err := p.productRepo.UpdateProduct(productModel)
	if err != nil {
		logrus.WithError(err).Error("Error updating product")
		return nil, err
	}

	productResponse := &response.ProductResponse{
		ProductID: updatedProduct.ID,
		Name:      updatedProduct.Name,
		Category:  updatedProduct.Category,
		Price:     updatedProduct.Price,
		Stock:     updatedProduct.Stock,
	}

	return productResponse, nil
}

func NewProductServiceImpl(productRepo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepo: productRepo}
}
