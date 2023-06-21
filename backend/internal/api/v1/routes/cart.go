package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/backend/internal/api/v1/handlers"
)

// CartRoutes adds new routes to the cart endpoint for a customer.
func CartRoutes(customer *gin.RouterGroup) {
	customer.POST("/cart", handlers.PostItemToCustomerCart)
	customer.DELETE("/cart/:item_id", handlers.DeleteItemFromCustomerCart)
	customer.DELETE("/cart", handlers.ClearCustomerCart)
	customer.GET("/cart", handlers.GetCustomerCart)
	customer.PUT("/cart/:item_id/:quantity", handlers.PutCartItemQuantity)
}
