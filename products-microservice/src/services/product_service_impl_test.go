package services

import (
	"testing"

	"github.com/dieg0code/products-microservice/src/json/request"
	"github.com/dieg0code/products-microservice/src/models"
	"github.com/dieg0code/products-microservice/src/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestProductServiceImpl(t *testing.T) {

	t.Run("CreateProduct_Success", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockReq := request.CreateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		mockModel := &models.Product{
			Name:     mockReq.Name,
			Category: mockReq.Category,
			Price:    mockReq.Price,
			Stock:    mockReq.Stock,
		}

		mockRepo.On("CreateProduct", mockModel).Return(&models.Product{
			Model:    gorm.Model{ID: 1},
			Name:     mockReq.Name,
			Category: mockReq.Category,
			Price:    mockReq.Price,
			Stock:    mockReq.Stock,
		}, nil)

		productID, err := productService.CreateProduct(&mockReq)

		assert.Nil(t, err, "Expected error to be nil")
		assert.Equal(t, uint(1), *productID, "Expected product ID to be 1")

		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateProduct_Error", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockReq := request.CreateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		mockModel := &models.Product{
			Name:     mockReq.Name,
			Category: mockReq.Category,
			Price:    mockReq.Price,
			Stock:    mockReq.Stock,
		}

		mockRepo.On("CreateProduct", mockModel).Return(&models.Product{}, assert.AnError)

		productID, err := productService.CreateProduct(&mockReq)

		assert.NotNil(t, err, "Expected error to be not nil")
		assert.Nil(t, productID, "Expected product ID to be nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteProduct_Success", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("DeleteProduct", uint(1)).Return(nil)

		err := productService.DeleteProduct(1)

		assert.Nil(t, err, "Expected error to be nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteProduct_Error", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("DeleteProduct", uint(1)).Return(assert.AnError)

		err := productService.DeleteProduct(1)

		assert.NotNil(t, err, "Expected error to be not nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllProducts_Success", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("GetAllProducts", 0, 10).Return([]models.Product{
			{
				Model:    gorm.Model{ID: 1},
				Name:     "Product 1",
				Category: "Category 1",
				Price:    1000,
				Stock:    10,
			},
			{
				Model:    gorm.Model{ID: 2},
				Name:     "Product 2",
				Category: "Category 2",
				Price:    2000,
				Stock:    20,
			},
		}, nil)

		products, err := productService.GetAllProducts(1, 10)

		assert.Nil(t, err, "Expected error to be nil")
		assert.Equal(t, 2, len(products), "Expected 2 products")

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllProducts_Error", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("GetAllProducts", 0, 10).Return([]models.Product{}, assert.AnError)

		products, err := productService.GetAllProducts(1, 10)

		assert.NotNil(t, err, "Expected error to be not nil")
		assert.Nil(t, products, "Expected products to be nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetByCategory_Success", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("GetByCategory", "Category 1").Return([]models.Product{
			{
				Model:    gorm.Model{ID: 1},
				Name:     "Product 1",
				Category: "Category 1",
				Price:    1000,
				Stock:    10,
			},
			{
				Model:    gorm.Model{ID: 2},
				Name:     "Product 2",
				Category: "Category 1",
				Price:    2000,
				Stock:    20,
			},
		}, nil)

		products, err := productService.GetByCategory("Category 1")

		assert.Nil(t, err, "Expected error to be nil")
		assert.Equal(t, 2, len(products), "Expected 2 products")

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetByCategory_Error", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("GetByCategory", "Category 1").Return([]models.Product{}, assert.AnError)

		products, err := productService.GetByCategory("Category 1")

		assert.NotNil(t, err, "Expected error to be not nil")
		assert.Nil(t, products, "Expected products to be nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetProductById_Success", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("GetProductById", uint(1)).Return(&models.Product{
			Model:    gorm.Model{ID: 1},
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}, nil)

		product, err := productService.GetProductById(1)

		assert.Nil(t, err, "Expected error to be nil")
		assert.NotNil(t, product, "Expected product to be not nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetProductById_Error", func(t *testing.T) {
		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockRepo.On("GetProductById", uint(1)).Return(&models.Product{}, assert.AnError)

		product, err := productService.GetProductById(1)

		assert.NotNil(t, err, "Expected error to be not nil")
		assert.Nil(t, product, "Expected product to be nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateProduct_Success", func(t *testing.T) {

		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockReq := &request.UpdateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		mockRepo.On("UpdateProduct", uint(1), mock.Anything).Return(&models.Product{
			Model:    gorm.Model{ID: 1},
			Name:     mockReq.Name,
			Category: mockReq.Category,
			Price:    mockReq.Price,
			Stock:    mockReq.Stock,
		}, nil)

		product, err := productService.UpdateProduct(1, mockReq)

		assert.Nil(t, err, "Expected error to be nil")
		assert.NotNil(t, product, "Expected product to be not nil")

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateProduct_Error", func(t *testing.T) {

		mockRepo := new(testutils.MockProductRepository)

		productService := NewProductServiceImpl(mockRepo)

		mockReq := &request.UpdateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		mockRepo.On("UpdateProduct", uint(1), mock.Anything).Return(&models.Product{}, assert.AnError)

		product, err := productService.UpdateProduct(1, mockReq)

		assert.NotNil(t, err, "Expected error to be not nil")
		assert.Nil(t, product, "Expected product to be nil")

		mockRepo.AssertExpectations(t)
	})

}
