package routes

import (
	"github.com/chee-zaram/wigit-webapp/backend/internal/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

// ProductsRoutes adds all routes for the product endpoint.
func ProductsRoutes(api *gin.RouterGroup) {
	api.GET("/products", handlers.GetProducts)
	api.GET("/products/:product_id", handlers.GetProductByID)
	api.GET("/products/categories/:category", handlers.GetProductsByCategory)
	api.GET("/products/search/:name", handlers.GetProductsByName)
}

// AdminProductsRoutes adds all routes for the admin products endpoint.
func AdminProductsRoutes(admin *gin.RouterGroup) {
	admin.POST("/products", handlers.AdminPostProduct)
	admin.DELETE("/products/:product_id", handlers.AdminDeleteProduct)
	admin.PUT("/products/:product_id", handlers.AdminPutProduct)
}
