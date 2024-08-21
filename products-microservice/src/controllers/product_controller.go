package controllers

import "github.com/gin-gonic/gin"

type ProductController interface {
	CreateProduct(c *gin.Context)
	GetProductById(c *gin.Context)
	GetAllProducts(c *gin.Context)
	GetByCategory(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}
