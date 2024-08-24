package repository

import (
	"testing"

	"github.com/dieg0code/products-microservice/src/models"
	"github.com/dieg0code/products-microservice/src/testutils"
	"github.com/stretchr/testify/assert"
)

func TestProductRespositoryImpl(t *testing.T) {

	t.Run("CheckProductExist_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		exists, err := repo.CheckProductExist(product.ID)

		assert.Nil(t, err, "Expected no error checking product existence")
		assert.True(t, exists, "Expected product to exist")
	})

	t.Run("CheckProductExist_Failure", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		exists, err := repo.CheckProductExist(1)

		assert.Nil(t, err, "Expected no error checking product existence")
		assert.False(t, exists, "Expected product to not exist")
	})

	t.Run("CreateProduct_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")
	})

	t.Run("CreateProduct_Failure", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		product, err = repo.CreateProduct(mockProduct)

		assert.NotNil(t, err, "Expected error creating product")
		assert.Nil(t, product, "Expected no product to be created")
	})

	t.Run("DeleteProduct_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		err = repo.DeleteProduct(product.ID)

		assert.Nil(t, err, "Expected no error deleting product")
	})

	t.Run("DeleteProduct_Failure_Not_Found", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		err := repo.DeleteProduct(1)

		assert.NotNil(t, err, "Expected error deleting product")
		assert.Equal(t, "product not found", err.Error(), "Expected error message to be 'product not found'")
	})

	t.Run("DeleteProduct_Failure_CheckProductExist_Error", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		err = repo.DeleteProduct(product.ID)

		assert.Nil(t, err, "Expected no error deleting product")

		err = repo.DeleteProduct(product.ID)

		assert.NotNil(t, err, "Expected error deleting product")
	})

	t.Run("GetAllProducts_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		products, err := repo.GetAllProducts(0, 10)

		assert.Nil(t, err, "Expected no error getting all products")
		assert.NotEmpty(t, products, "Expected products to be returned")
		assert.Equal(t, 1, len(products), "Expected 1 product to be returned")
		assert.Equal(t, product.ID, products[0].ID, "Expected product ID to be the same")
	})

	t.Run("GetAllProducts_Failure", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		products, err := repo.GetAllProducts(0, 10)

		assert.Nil(t, err, "Expected no error getting all products")
		assert.Empty(t, products, "Expected no products to be returned")
	})

	t.Run("GetByCategory_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		products, err := repo.GetByCategory("Test Category")

		assert.Nil(t, err, "Expected no error getting products by category")
		assert.NotEmpty(t, products, "Expected products to be returned")
		assert.Equal(t, 1, len(products), "Expected 1 product to be returned")
		assert.Equal(t, product.ID, products[0].ID, "Expected product ID to be the same")
	})

	t.Run("GetByCategory_Failure", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		products, err := repo.GetByCategory("Test Category")

		assert.Nil(t, err, "Expected no error getting products by category")
		assert.Empty(t, products, "Expected no products to be returned")
	})

	t.Run("GetProductById_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		product, err = repo.GetProductById(product.ID)

		assert.Nil(t, err, "Expected no error getting product by id")
		assert.NotNil(t, product, "Expected product to be returned")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")
	})

	t.Run("GetProductById_Failure_Not_Found", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		product, err := repo.GetProductById(1)

		assert.NotNil(t, err, "Expected error getting product by id")
		assert.Nil(t, product, "Expected no product to be returned")
	})

	t.Run("UpdateProduct_Success", func(t *testing.T) {
		db := testutils.SetupTestDB(&models.Product{})
		defer func() {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				t.Error("Error closing database connection")
			}
		}()

		repo := NewPorductRespositoryImpl(db)

		mockProduct := &models.Product{
			Name:     "Test Product",
			Category: "Test Category",
			Price:    1000,
			Stock:    10,
		}

		product, err := repo.CreateProduct(mockProduct)

		assert.Nil(t, err, "Expected no error creating product")
		assert.NotNil(t, product, "Expected product to be created")
		assert.NotEqual(t, uint(0), product.ID, "Expected product ID to be set")
		assert.Equal(t, mockProduct.Name, product.Name, "Expected product name to be the same")

		product.Name = "Updated Product"

		product, err = repo.UpdateProduct(product.ID, product)

		assert.Nil(t, err, "Expected no error updating product")
		assert.NotNil(t, product, "Expected product to be updated")
		assert.Equal(t, "Updated Product", product.Name, "Expected product name to be updated")
	})

}
