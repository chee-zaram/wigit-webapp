package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// AdminGetBookings retrieves all bookings in the database.
func AdminGetBookings(ctx *gin.Context) {
	var bookings []models.Booking

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Find(&bookings).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	if bookings == nil {
		bookings = []models.Booking{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}
