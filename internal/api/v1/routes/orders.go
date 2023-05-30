package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// OrdersRoutes adds routes for the customer in the orders endpoint.
func OrdersRoutes(customer *gin.RouterGroup) {
	customer.POST("/orders", handlers.CustomerPostOrders)
	customer.GET("/orders", handlers.CustomerGetOrders)
	customer.GET("/orders/:order_id", handlers.CustomerGetOrderByID)
	customer.GET("/orders/status/:status", handlers.CustomerGetOrdersByStatus)
}

// AdminOrdersRoutes adds routes for the admin in the orders endpoint.
func AdminOrdersRoutes(admin *gin.RouterGroup) {
	// Only the status field can be updated.
	admin.PUT("/orders/:order_id/:status", handlers.AdminPutOrders)
	admin.GET("/orders", handlers.AdminGetOrders)
	admin.GET("/orders/:order_id", handlers.AdminGetOrderByID)
	admin.GET("/orders/status/:status", handlers.AdminGetOrdersByStatus)
}
