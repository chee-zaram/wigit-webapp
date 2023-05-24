package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// ProductsRoutes adds all routes for the product endpoint.
func ProductsRoutes(api *gin.RouterGroup) {
	api.GET("/products", handlers.GetProducts)
	api.GET("/products/:product_id", handlers.GetProductByID)
	api.GET("/products/categories/:category", handlers.GetProductByCategory)
}

// AdminProductsRoutes adds all routes for the admin products endpoint.
func AdminProductsRoutes(admin *gin.RouterGroup) {
	admin.POST("/products", handlers.AdminPostProducts)
	admin.DELETE("/products/:product_id", handlers.AdminDeleteProducts)
	admin.PUT("/products/:product_id", handlers.AdminPutProducts)
}
