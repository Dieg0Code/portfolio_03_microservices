package testutils

import (
	"github.com/dieg0code/products-microservice/src/models"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*models.Product), args.Error(1)
}
func (m *MockProductRepository) GetProductById(ProductID uint) (*models.Product, error) {
	args := m.Called(ProductID)
	return args.Get(0).(*models.Product), args.Error(1)
}
func (m *MockProductRepository) GetAllProducts(offset int, pageSize int) ([]models.Product, error) {
	args := m.Called(offset, pageSize)
	return args.Get(0).([]models.Product), args.Error(1)
}
func (m *MockProductRepository) GetByCategory(category string) ([]models.Product, error) {
	args := m.Called(category)
	return args.Get(0).([]models.Product), args.Error(1)
}
func (m *MockProductRepository) UpdateProduct(productID uint, product *models.Product) (*models.Product, error) {
	args := m.Called(productID, product)
	return args.Get(0).(*models.Product), args.Error(1)
}
func (m *MockProductRepository) DeleteProduct(ProductID uint) error {
	args := m.Called(ProductID)
	return args.Error(0)
}
func (m *MockProductRepository) CheckProductExist(ProductID uint) (bool, error) {
	args := m.Called(ProductID)
	return args.Bool(0), args.Error(1)
}
