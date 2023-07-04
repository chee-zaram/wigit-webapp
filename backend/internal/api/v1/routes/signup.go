package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-ng/webapp/backend/internal/api/v1/handlers"
)

// SignUpRoutes adds the routes all routes configured for the signup endpoint.
func SignUpRoutes(api *gin.RouterGroup) {
	api.POST("/signup", handlers.SignUp)
}
