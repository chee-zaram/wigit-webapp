package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/handlers"
)

// ServicesRoutes adds public routes to the services endpoint.
func ServicesRoutes(api *gin.RouterGroup) {
	api.GET("/services", handlers.GetServices)
	api.GET("/services/:service_id", handlers.GetServiceByID)
}
