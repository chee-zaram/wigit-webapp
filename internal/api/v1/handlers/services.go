package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// GetServices retrieves a list of all available services.
func GetServices(ctx *gin.Context) {
	var services []models.Service

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Find(&services).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	if services == nil {
		services = []models.Service{}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": services})
}

// GetServiceByID retrieves a service from the database based on given service id.
func GetServiceByID(ctx *gin.Context) {
	id := ctx.Param("service_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidServiceID.Error()})
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidServiceID.Error()})
		return
	}

	service := new(models.Service)
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(service, "id = ?", id).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidServiceID.Error()})
		return
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": service})
}
