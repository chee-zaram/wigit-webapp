package routes

import (
	"github.com/chee-zaram/wigit-webapp/backend/internal/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

// SlotsRoutes adds new routes to the slots endpoint.
func SlotsRoutes(api *gin.RouterGroup) {
	api.GET("/slots", handlers.GetSlots)
}

// AdminSlotsRoutes adds new admin routes for the slots endpoint.
func AdminSlotsRoutes(admin *gin.RouterGroup) {
	admin.POST("/slots", handlers.AdminPostSlots)
	admin.DELETE("/slots/:slot_id", handlers.AdminDeleteSlots)
}
