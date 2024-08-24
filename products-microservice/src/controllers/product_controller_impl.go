package controllers

import (
	"strconv"

	"github.com/dieg0code/products-microservice/src/json/request"
	"github.com/dieg0code/products-microservice/src/json/response"
	"github.com/dieg0code/products-microservice/src/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
	validate       *validator.Validate
}

// CreateProduct implements ProductController.
func (p *ProductControllerImpl) CreateProduct(c *gin.Context) {

	createProductRequest := &request.CreateProductRequest{}

	err := c.ShouldBindJSON(createProductRequest)
	if err != nil {
		logrus.WithError(err).Error("Error binding request")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid request body",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	err = p.validate.Struct(createProductRequest)
	if err != nil {
		logrus.WithError(err).Error("Error validating request")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid request body",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	productID, err := p.ProductService.CreateProduct(createProductRequest)
	if err != nil {
		logrus.WithError(err).Error("Error creating product")
		errRes := response.BaseResponse{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    "Error creating product",
			Data:   nil,
		}

		c.JSON(500, errRes)
		return
	}

	res := response.BaseResponse{
		Code:   201,
		Status: "Created",
		Msg:    "Product created successfully",
		Data:   productID,
	}

	c.JSON(201, res)
}

// DeleteProduct implements ProductController.
func (p *ProductControllerImpl) DeleteProduct(c *gin.Context) {
	productID := c.Param("productID")

	productIDUint, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		logrus.WithError(err).Error("Error parsing productID")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid productID",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	id := uint(productIDUint)

	err = p.ProductService.DeleteProduct(id)
	if err != nil {
		logrus.WithError(err).Error("Error deleting product")
		errRes := response.BaseResponse{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    "Error deleting product",
			Data:   nil,
		}

		c.JSON(500, errRes)
		return
	}

	res := response.BaseResponse{
		Code:   200,
		Status: "OK",
		Msg:    "Product deleted successfully",
		Data:   nil,
	}

	c.JSON(200, res)
}

// GetAllProducts implements ProductController.
func (p *ProductControllerImpl) GetAllProducts(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		logrus.WithError(err).Error("Error parsing page")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid page",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		logrus.WithError(err).Error("Error parsing pageSize")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid pageSize",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	products, err := p.ProductService.GetAllProducts(pageInt, pageSizeInt)
	if err != nil {
		logrus.WithError(err).Error("Error getting all products")
		errRes := response.BaseResponse{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    "Error getting all products",
			Data:   nil,
		}

		c.JSON(500, errRes)
		return
	}

	res := response.BaseResponse{
		Code:   200,
		Status: "OK",
		Msg:    "Products retrieved successfully",
		Data:   products,
	}

	c.JSON(200, res)
}

// GetByCategory implements ProductController.
func (p *ProductControllerImpl) GetByCategory(c *gin.Context) {

	category := c.Param("category")

	products, err := p.ProductService.GetByCategory(category)
	if err != nil {
		logrus.WithError(err).Error("Error getting products by category")
		errRes := response.BaseResponse{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    "Error getting products by category",
			Data:   nil,
		}

		c.JSON(500, errRes)
		return
	}

	res := response.BaseResponse{
		Code:   200,
		Status: "OK",
		Msg:    "Products retrieved successfully",
		Data:   products,
	}

	c.JSON(200, res)
}

// GetProductById implements ProductController.
func (p *ProductControllerImpl) GetProductById(c *gin.Context) {

	productID := c.Param("productID")

	productIDUint, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		logrus.WithError(err).Error("Error parsing productID")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid productID",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	id := uint(productIDUint)

	product, err := p.ProductService.GetProductById(id)
	if err != nil {
		logrus.WithError(err).Error("Error getting product by ID")
		errRes := response.BaseResponse{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    "Error getting product by ID",
			Data:   nil,
		}

		c.JSON(500, errRes)
		return
	}

	res := response.BaseResponse{
		Code:   200,
		Status: "OK",
		Msg:    "Product retrieved successfully",
		Data:   product,
	}

	c.JSON(200, res)
}

// UpdateProduct implements ProductController.
func (p *ProductControllerImpl) UpdateProduct(c *gin.Context) {

	productID := c.Param("productID")

	productIDUint, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		logrus.WithError(err).Error("Error parsing productID")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid productID",
			Data:   nil,
		}

		c.JSON(400, errRes)
	}

	updateProductRequest := &request.UpdateProductRequest{}

	err = c.ShouldBindJSON(updateProductRequest)
	if err != nil {
		logrus.WithError(err).Error("Error binding request")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid request body",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	err = p.validate.Struct(updateProductRequest)
	if err != nil {
		logrus.WithError(err).Error("Error validating request")
		errRes := response.BaseResponse{
			Code:   400,
			Status: "Bad Request",
			Msg:    "Invalid request body",
			Data:   nil,
		}

		c.JSON(400, errRes)
		return
	}

	product, err := p.ProductService.UpdateProduct(uint(productIDUint), updateProductRequest)
	if err != nil {
		logrus.WithError(err).Error("Error updating product")
		errRes := response.BaseResponse{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    "Error updating product",
			Data:   nil,
		}

		c.JSON(500, errRes)
		return
	}

	res := response.BaseResponse{
		Code:   200,
		Status: "OK",
		Msg:    "Product updated successfully",
		Data:   product,
	}

	c.JSON(200, res)
}

func NewProductControllerImpl(productService services.ProductService, validate *validator.Validate) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
		validate:       validate,
	}
}
