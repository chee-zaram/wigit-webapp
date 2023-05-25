package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// CartRoutes adds new routes to the cart endpoint for a customer.
func CartRoutes(customer *gin.RouterGroup) {
	customer.POST("/cart", handlers.CustomerPostToCart)
	customer.DELETE("/cart/:item_id", handlers.CustomerDeleteFromCart)
	customer.DELETE("/cart", handlers.CustomerClearCart)
	customer.GET("/cart", handlers.CustomerGetCart)
}
