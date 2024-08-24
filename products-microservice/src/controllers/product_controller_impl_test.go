package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dieg0code/products-microservice/src/json/request"
	"github.com/dieg0code/products-microservice/src/json/response"
	"github.com/dieg0code/products-microservice/src/testutils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestProductControllerImpl(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("CreateProduct_Success", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.POST("/products", controller.CreateProduct)

		reqBody := &request.CreateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		productID := uint(1)
		mockService.On("CreateProduct", reqBody).Return(&productID, nil)

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err, "Expected no error marshalling request body")

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		assert.Nil(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code, "Expected status code 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 201, response.Code, "Expected response code 200")
		assert.Equal(t, "Created", response.Status, "Expected response status OK")

		mockService.AssertExpectations(t)
	})

	t.Run("CreateProduct_BadRequest", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.POST("/products", controller.CreateProduct)

		reqBody := &request.CreateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			// Price:    1000, not provided
			// Stock:    10, not provided
			// Shoud fail validation
		}

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err, "Expected no error marshalling request body")

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status code 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 400, response.Code, "Expected response code 400")
		assert.Equal(t, "Bad Request", response.Status, "Expected response status Bad Request")
		assert.Equal(t, "Invalid request body", response.Msg, "Expected response message Invalid request body")
	})

	t.Run("CreateProduct_InternalServerError", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.POST("/products", controller.CreateProduct)

		reqBody := &request.CreateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		productID := uint(1)
		mockService.On("CreateProduct", reqBody).Return(&productID, assert.AnError)

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err, "Expected no error marshalling request body")

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		assert.Nil(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status code 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 500, response.Code, "Expected response code 500")
		assert.Equal(t, "Internal Server Error", response.Status, "Expected response status Internal Server Error")
		assert.Equal(t, "Error creating product", response.Msg, "Expected response message Error creating product")

		mockService.AssertExpectations(t)
	})

	t.Run("DeleteProduct_Success", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.DELETE("/products/:productID", controller.DeleteProduct)

		productID := uint(1)
		mockService.On("DeleteProduct", productID).Return(nil)

		req, err := http.NewRequest(http.MethodDelete, "/products/1", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 200, response.Code, "Expected response code 200")
		assert.Equal(t, "OK", response.Status, "Expected response status OK")

		mockService.AssertExpectations(t)
	})

	t.Run("DeleteProduct_BadRequest", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.DELETE("/products/:productID", controller.DeleteProduct)

		req, err := http.NewRequest(http.MethodDelete, "/products/abc", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status code 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 400, response.Code, "Expected response code 400")
		assert.Equal(t, "Bad Request", response.Status, "Expected response status Bad Request")
		assert.Equal(t, "Invalid productID", response.Msg, "Expected response message Invalid productID")

		mockService.AssertExpectations(t)
	})

	t.Run("DeleteProduct_InternalServerError", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.DELETE("/products/:productID", controller.DeleteProduct)

		productID := uint(1)
		mockService.On("DeleteProduct", productID).Return(assert.AnError)

		req, err := http.NewRequest(http.MethodDelete, "/products/1", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status code 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 500, response.Code, "Expected response code 500")
		assert.Equal(t, "Internal Server Error", response.Status, "Expected response status Internal Server Error")
		assert.Equal(t, "Error deleting product", response.Msg, "Expected response message Error deleting product")

		mockService.AssertExpectations(t)
	})

	t.Run("GetAllProducts_Success", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products", controller.GetAllProducts)

		page := 1
		pageSize := 10

		mockService.On("GetAllProducts", page, pageSize).Return([]response.ProductResponse{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/products?page=1&pageSize=10", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 200, response.Code, "Expected response code 200")

		mockService.AssertExpectations(t)

	})

	t.Run("GetAllProducts_BadRequest", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products", controller.GetAllProducts)

		req, err := http.NewRequest(http.MethodGet, "/products?page=abc&pageSize=10", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status code 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 400, response.Code, "Expected response code 400")
		assert.Equal(t, "Bad Request", response.Status, "Expected response status Bad Request")
		assert.Equal(t, "Invalid page", response.Msg, "Expected response message Invalid page")

		mockService.AssertExpectations(t)
	})

	t.Run("GetAllProducts_InternalServerError", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products", controller.GetAllProducts)

		page := 1
		pageSize := 10

		mockService.On("GetAllProducts", page, pageSize).Return([]response.ProductResponse{}, assert.AnError)

		req, err := http.NewRequest(http.MethodGet, "/products?page=1&pageSize=10", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status code 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 500, response.Code, "Expected response code 500")
		assert.Equal(t, "Internal Server Error", response.Status, "Expected response status Internal Server Error")
		assert.Equal(t, "Error getting all products", response.Msg, "Expected response message Error getting all products")

		mockService.AssertExpectations(t)
	})

	t.Run("GetByCategory_Success", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products/category/:category", controller.GetByCategory)

		category := "Category 1"

		mockService.On("GetByCategory", category).Return([]response.ProductResponse{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/products/category/Category%201", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 200, response.Code, "Expected response code 200")

		mockService.AssertExpectations(t)
	})

	t.Run("GetByCategory_InternalServerError", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products/category/:category", controller.GetByCategory)

		category := "Category 1"

		mockService.On("GetByCategory", category).Return([]response.ProductResponse{}, assert.AnError)

		req, err := http.NewRequest(http.MethodGet, "/products/category/Category%201", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status code 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 500, response.Code, "Expected response code 200")

		mockService.AssertExpectations(t)
	})

	t.Run("GetProductById_Success", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products/:productID", controller.GetProductById)

		productID := uint(1)

		mockService.On("GetProductById", productID).Return(&response.ProductResponse{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/products/1", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 200, response.Code, "Expected response code 200")

		mockService.AssertExpectations(t)

	})

	t.Run("GetProductById_BadRequest", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products/:productID", controller.GetProductById)

		req, err := http.NewRequest(http.MethodGet, "/products/abc", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status code 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 400, response.Code, "Expected response code 400")
		assert.Equal(t, "Bad Request", response.Status, "Expected response status Bad Request")
		assert.Equal(t, "Invalid productID", response.Msg, "Expected response message Invalid productID")

		mockService.AssertExpectations(t)
	})

	t.Run("GetProductById_InternalServerError", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.GET("/products/:productID", controller.GetProductById)

		productID := uint(1)

		mockService.On("GetProductById", productID).Return(&response.ProductResponse{}, assert.AnError)

		req, err := http.NewRequest(http.MethodGet, "/products/1", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status code 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 500, response.Code, "Expected response code 500")
		assert.Equal(t, "Internal Server Error", response.Status, "Expected response status Internal Server Error")
		assert.Equal(t, "Error getting product by ID", response.Msg, "Expected response message Error getting product by ID")

		mockService.AssertExpectations(t)
	})

	t.Run("UpdateProduct_Success", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.PUT("/products/:productID", controller.UpdateProduct)

		productID := uint(1)

		reqBody := &request.UpdateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		mockService.On("UpdateProduct", productID, reqBody).Return(&response.ProductResponse{}, nil)

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err, "Expected no error marshalling request body")

		req, err := http.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(body))
		assert.Nil(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 200, response.Code, "Expected response code 200")

		mockService.AssertExpectations(t)
	})

	t.Run("UpdateProduct_BadRequest", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.PUT("/products/:productID", controller.UpdateProduct)

		reqBody := &request.UpdateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			// Price:    1000, not provided
			// Stock:    10, not provided
			// Shoud fail validation
		}

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err, "Expected no error marshalling request body")

		req, err := http.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(body))
		assert.Nil(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status code 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 400, response.Code, "Expected response code 400")
		assert.Equal(t, "Bad Request", response.Status, "Expected response status Bad Request")
		assert.Equal(t, "Invalid request body", response.Msg, "Expected response message Invalid request body")

		mockService.AssertExpectations(t)
	})

	t.Run("UpdateProduct_InternalServerError", func(t *testing.T) {
		mockService := new(testutils.MockProductService)
		validator := validator.New()
		controller := NewProductControllerImpl(mockService, validator)

		router := gin.Default()
		router.PUT("/products/:productID", controller.UpdateProduct)

		productID := uint(1)

		reqBody := &request.UpdateProductRequest{
			Name:     "Product 1",
			Category: "Category 1",
			Price:    1000,
			Stock:    10,
		}

		mockService.On("UpdateProduct", productID, reqBody).Return(&response.ProductResponse{}, assert.AnError)

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err, "Expected no error marshalling request body")

		req, err := http.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(body))
		assert.Nil(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status code 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Nil(t, err, "Expected no error unmarshalling response body")
		assert.Equal(t, 500, response.Code, "Expected response code 500")
		assert.Equal(t, "Internal Server Error", response.Status, "Expected response status Internal Server Error")
		assert.Equal(t, "Error updating product", response.Msg, "Expected response message Error updating product")

		mockService.AssertExpectations(t)
	})

}
