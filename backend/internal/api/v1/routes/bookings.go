package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/backend/internal/api/v1/handlers"
)

// AdminBookingsRoutes adds all the routes for the bookings endpoint used by the admin.
func AdminBookingsRoutes(admin *gin.RouterGroup) {
	admin.GET("/bookings", handlers.AdminGetBookings)
	admin.GET("/bookings/:booking_id", handlers.AdminGetBooking)
	admin.PUT("/bookings/:booking_id/:status", handlers.AdminPutBooking)
}

// BookingsRoutes adds new routes for the bookings endpoint for customers.
func BookingsRoutes(customer *gin.RouterGroup) {
	customer.POST("/bookings", handlers.CustomerPostBooking)
	customer.GET("/bookings/:booking_id", handlers.CustomerGetBooking)
	customer.GET("/bookings", handlers.CustomerGetBookings)
}
