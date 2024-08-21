package router

import (
	"github.com/dieg0code/products-microservice/src/controllers"
	"github.com/dieg0code/products-microservice/src/db"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ProductController controllers.ProductController
}

func NewRouter(productController controllers.ProductController) *Router {
	return &Router{
		ProductController: productController,
	}
}

func (r *Router) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to Products Microservice",
		})
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Products Microservice is healthy",
		})
	})

	router.GET("/ready", func(ctx *gin.Context) {

		dbConn := db.DatabaseConnection()
		err := db.CheckDatabaseConnection(dbConn)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Products Microservice is not ready",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Products Microservice is ready",
		})
	})

	baseRoute := router.Group("/api/v1")
	{
		productRoute := baseRoute.Group("/products")
		{
			productRoute.POST("", r.ProductController.CreateProduct)
			productRoute.GET("/:productID", r.ProductController.GetProductById)
			productRoute.GET("", r.ProductController.GetAllProducts)
			productRoute.GET("/category/:category", r.ProductController.GetByCategory)
			productRoute.PUT("/:productID", r.ProductController.UpdateProduct)
			productRoute.DELETE("/:productID", r.ProductController.DeleteProduct)
		}
	}

	return router
}
