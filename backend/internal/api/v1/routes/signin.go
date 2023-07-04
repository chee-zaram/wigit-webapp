package routes

import (
	"github.com/chee-zaram/wigit-webapp/backend/internal/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

// SignInRoutes add all routes in the signin endpoint.
func SignInRoutes(api *gin.RouterGroup) {
	api.POST("/signin", handlers.SignIn)
}
