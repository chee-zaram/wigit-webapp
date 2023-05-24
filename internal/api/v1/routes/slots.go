package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// SlotsRoutes adds new routes to the slots endpoint.
func SlotsRoutes(api *gin.RouterGroup) {
	api.GET("/slots", handlers.GetSlots)
}

// AdminSlotsRoutes adds new admin routes for the slots endpoint.
func AdminSlotsRoutes(admin *gin.RouterGroup) {
	admin.POST("/slots", handlers.AdminPostSlots)
}
