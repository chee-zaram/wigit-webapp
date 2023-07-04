package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-ng/webapp/backend/internal/api/v1/handlers"
)

// ResetPasswordRoutes adds the routes for the reset-password endpoint.
func ResetPasswordRoutes(api *gin.RouterGroup) {
	api.POST("/reset_password", handlers.PostResetPassword)
	api.PUT("/reset_password", handlers.PutResetPassword)
}
