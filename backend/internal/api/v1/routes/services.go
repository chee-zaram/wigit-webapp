package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-ng/webapp/backend/internal/api/v1/handlers"
	"github.com/wigit-ng/webapp/backend/internal/api/v1/middlewares"
)

// ServicesRoutes adds public routes to the services endpoint.
func ServicesRoutes(api *gin.RouterGroup) {
	api.GET("/services", middlewares.Redis, handlers.GetServices)
	api.GET("/services/:service_id", middlewares.Redis, handlers.GetServiceByID)
	api.GET("/services/trending", middlewares.Redis, handlers.GetTrendingServices)
}

// AdminServicesRoutes add protected routes for the services endpoint.
func AdminServicesRoutes(admin *gin.RouterGroup) {
	admin.POST("/services", handlers.AdminPostService)
	admin.DELETE("/services/:service_id", handlers.AdminDeleteService)
	admin.PUT("/services/:service_id", handlers.AdminPutService)
}
