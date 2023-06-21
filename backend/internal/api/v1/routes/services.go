package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/backend/internal/api/v1/handlers"
)

// ServicesRoutes adds public routes to the services endpoint.
func ServicesRoutes(api *gin.RouterGroup) {
	api.GET("/services", handlers.GetServices)
	api.GET("/services/:service_id", handlers.GetServiceByID)
	api.GET("/services/trending", handlers.GetTrendingServices)
}

// AdminServicesRoutes add protected routes for the services endpoint.
func AdminServicesRoutes(admin *gin.RouterGroup) {
	admin.POST("/services", handlers.AdminPostService)
	admin.DELETE("/services/:service_id", handlers.AdminDeleteService)
	admin.PUT("/services/:service_id", handlers.AdminPutService)
}
