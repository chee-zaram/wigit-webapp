package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// UsersRoutes adds adds new routes for deleting user and updating user information.
func UsersRoutes(customer *gin.RouterGroup) {
	customer.DELETE("/users/:user_id", handlers.CustomerDeleteUser)
	customer.PUT("/users/:user_id", handlers.CustomerPutUser)
}

// AdminUsersRoutes adds new routes for retrieving a user's order info.
func AdminUsersRoutes(admin *gin.RouterGroup) {
	admin.GET("/users/:email/orders_bookings", handlers.AdminGetUserOrdersBookings)
}

// SuperAdminUsersRoutes adds new routes for manipulating user information.
func SuperAdminUsersRoutes(superAdmin *gin.RouterGroup) {
	superAdmin.PUT("/users/:email/:new_role", handlers.SuperAdminUpdateRole)
}
