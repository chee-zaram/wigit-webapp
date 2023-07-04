package routes

import (
	"github.com/chee-zaram/wigit-webapp/backend/internal/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

// SignUpRoutes adds the routes all routes configured for the signup endpoint.
func SignUpRoutes(api *gin.RouterGroup) {
	api.POST("/signup", handlers.SignUp)
}
