package repository

import (
	"errors"

	"github.com/dieg0code/products-microservice/src/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

// CheckProductExist implements ProductRepository.
func (p *ProductRepositoryImpl) CheckProductExist(ProductID uint) (bool, error) {
	var exists int64

	res := p.db.Model(&models.Product{}).Where(IdPlaceholder, ProductID).Count(&exists)

	if res.Error != nil {
		logrus.WithError(res.Error).Error("Error checking product existence")
		return false, res.Error
	}

	return exists > 0, nil
}

// CreateProduct implements ProductRepository.
func (p *ProductRepositoryImpl) CreateProduct(product *models.Product) (*models.Product, error) {

	result := p.db.Create(product)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error creating product")
		return nil, result.Error
	}

	return product, result.Error
}

// DeleteProduct implements ProductRepository.
func (p *ProductRepositoryImpl) DeleteProduct(ProductID uint) error {
	exists, err := p.CheckProductExist(ProductID)

	if err != nil {
		logrus.WithError(err).Error("Error checking product existence")
	}

	if !exists {
		logrus.Error("Product not found")
		return errors.New("product not found")
	}

	result := p.db.Delete(&models.Product{}, ProductID)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error deleting product")
		return result.Error
	}

	return nil
}

// GetAllProducts implements ProductRepository.
func (p *ProductRepositoryImpl) GetAllProducts(offset int, pageSize int) ([]models.Product, error) {
	var products []models.Product

	res := p.db.Offset(offset).Limit(pageSize).Find(&products)
	if res.Error != nil {
		logrus.WithError(res.Error).Error("Error getting all products")
		return nil, res.Error
	}

	return products, nil
}

// GetByCategory implements ProductRepository.
func (p *ProductRepositoryImpl) GetByCategory(category string) ([]models.Product, error) {

	var products []models.Product

	res := p.db.Where(CategoryPlaceholder, category).Find(&products)
	if res.Error != nil {
		logrus.WithError(res.Error).Error("Error getting products by category")
		return nil, res.Error
	}

	return products, nil
}

// GetProductById implements ProductRepository.
func (p *ProductRepositoryImpl) GetProductById(ProductID uint) (*models.Product, error) {
	exists, err := p.CheckProductExist(ProductID)
	if err != nil {
		logrus.WithError(err).Error("Error checking product existence")
		return nil, err
	}

	if !exists {
		logrus.Error("Product not found")
		return nil, errors.New("product not found")
	}

	var product models.Product

	res := p.db.First(&product, ProductID)
	if res.Error != nil {
		logrus.WithError(res.Error).Error("Error getting product by id")
		return nil, res.Error
	}

	return &product, nil
}

// UpdateProduct implements ProductRepository.
func (p *ProductRepositoryImpl) UpdateProduct(prodctID uint, product *models.Product) (*models.Product, error) {
	exists, err := p.CheckProductExist(product.ID)
	if err != nil {
		logrus.WithError(err).Error("Error checking product existence")
		return nil, err
	}

	if !exists {
		logrus.WithField("product_id", product.ID).Errorf("Product with id %d not found", product.ID)
		return nil, errors.New("product not found")
	}

	result := p.db.Where(IdPlaceholder, product.ID).Updates(product)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error updating product")
		return nil, result.Error
	}

	return product, nil
}

func NewPorductRespositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}
