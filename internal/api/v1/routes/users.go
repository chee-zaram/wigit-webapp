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
