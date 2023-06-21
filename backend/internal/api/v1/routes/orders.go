package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/backend/internal/api/v1/handlers"
)

// OrdersRoutes adds routes for the customer in the orders endpoint.
func OrdersRoutes(customer *gin.RouterGroup) {
	customer.POST("/orders", handlers.PostCustomerOrder)
	customer.GET("/orders", handlers.GetCustomerOrders)
	customer.GET("/orders/:order_id", handlers.GetCustomerOrder)
	customer.GET("/orders/status/:status", handlers.GetOrdersByStatus)
}

// AdminOrdersRoutes adds routes for the admin in the orders endpoint.
func AdminOrdersRoutes(admin *gin.RouterGroup) {
	// Only the status field can be updated.
	admin.PUT("/orders/:order_id/:status", handlers.AdminPutOrders)
	admin.GET("/orders", handlers.AdminGetOrders)
	admin.GET("/orders/:order_id", handlers.AdminGetOrder)
	admin.GET("/orders/status/:status", handlers.AdminGetOrdersByStatus)
}
