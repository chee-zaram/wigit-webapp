package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// GetSlots retrieves a list of all available slots.
func GetSlots(ctx *gin.Context) {
	var slots []models.Slot

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Where("is_free = ?", true).Where("time > ?", time.Now()).Find(&slots).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	if slots == nil {
		slots = []models.Slot{}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": slots})
}
