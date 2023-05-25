package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// OrdersRoutes adds routes for the customer in the orders endpoint.
func OrdersRoutes(customer *gin.RouterGroup) {
	customer.POST("/orders", handlers.PostOrders)
}

// AdminOrdersRoutes adds routes for the admin in the orders endpoint.
func AdminOrdersRoutes(admin *gin.RouterGroup) {}
