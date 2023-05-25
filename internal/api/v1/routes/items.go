package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// ItemsRoutes adds new routes to the items endpoint for a customer.
func ItemsRoutes(customer *gin.RouterGroup) {
	customer.POST("/cart", handlers.CustomerPostItem)
	customer.DELETE("/cart/:item_id", handlers.CustomerDeleteItem)
	customer.DELETE("/cart", handlers.CustomerClearCart)
	customer.GET("/cart", handlers.CustomerGetCart)
}
