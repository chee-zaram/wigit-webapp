package routes

import (
	"github.com/chee-zaram/wigit-webapp/backend/internal/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

// UsersRoutes adds adds new routes for deleting user and updating user information.
func UsersRoutes(customer *gin.RouterGroup) {
	customer.DELETE("/users/:email", handlers.CustomerDeleteUser)
	customer.PUT("/users/:email", handlers.CustomerPutUser)
}

// AdminUsersRoutes adds new routes for retrieving a user's order info.
func AdminUsersRoutes(admin *gin.RouterGroup) {
	admin.GET("/users/:email/orders_bookings", handlers.AdminGetUserOrdersBookings)
}

// SuperAdminUsersRoutes adds new routes for manipulating user information.
func SuperAdminUsersRoutes(superAdmin *gin.RouterGroup) {
	superAdmin.PUT("/users/:email/:new_role", handlers.SuperAdminUpdateRole)
	superAdmin.DELETE("/users/:email", handlers.SuperAdminDeleteUser)
	superAdmin.GET("/users/admins", handlers.SuperAdminGetAdmins)
	superAdmin.GET("/users/customers", handlers.SuperAdminGetCustomers)
	superAdmin.GET("/users/:email", handlers.SuperAdminGetUser)
}
