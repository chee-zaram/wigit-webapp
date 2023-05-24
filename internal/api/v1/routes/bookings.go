package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// AdminBookingsRoutes adds all the routes for the bookings endpoint used by the admin.
func AdminBookingsRoutes(admin *gin.RouterGroup) {
	admin.GET("/bookings", handlers.AdminGetBookings)
}
