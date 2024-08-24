package testutils

import (
	"github.com/dieg0code/products-microservice/src/json/request"
	"github.com/dieg0code/products-microservice/src/json/response"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) CreateProduct(product *request.CreateProductRequest) (*uint, error) {
	args := m.Called(product)
	return args.Get(0).(*uint), args.Error(1)
}
func (m *MockProductService) GetProductById(productID uint) (*response.ProductResponse, error) {
	args := m.Called(productID)
	return args.Get(0).(*response.ProductResponse), args.Error(1)
}
func (m *MockProductService) GetAllProducts(page int, pageSize int) ([]response.ProductResponse, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]response.ProductResponse), args.Error(1)
}
func (m *MockProductService) GetByCategory(category string) ([]response.ProductResponse, error) {
	args := m.Called(category)
	return args.Get(0).([]response.ProductResponse), args.Error(1)
}
func (m *MockProductService) UpdateProduct(productID uint, product *request.UpdateProductRequest) (*response.ProductResponse, error) {
	args := m.Called(productID, product)
	return args.Get(0).(*response.ProductResponse), args.Error(1)
}
func (m *MockProductService) DeleteProduct(ProductID uint) error {
	args := m.Called(ProductID)
	return args.Error(0)
}
