package handlers

import (
	"errors"
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
		return tx.Where("is_free = ?", true).Where("date_time > ?", time.Now()).Find(&slots).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	if slots == nil {
		slots = []models.Slot{}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": slots})
}

// AdminPostSlots adds a new slot to the database.
func AdminPostSlots(ctx *gin.Context) {
	_slot := new(models.Slot)

	if err := ctx.ShouldBindJSON(_slot); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validatePostSlotsData(_slot); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Create(_slot).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	slot, err := getSlotFromDB(*_slot.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Slot created successfully",
		"data": slot,
	})
}

// validatePostSlotsData checks the validity of the json fields provided during
// a post request by an admin.
func validatePostSlotsData(slot *models.Slot) error {
	if slot.DateTime.Compare(time.Now()) < 1 {
		return errors.New("Date and time for slot must be in the future")
	}
	return nil
}

// getSlotFromDB retrieves a slot with id from the database.
func getSlotFromDB(id string) (*models.Slot, error) {
	slot := new(models.Slot)

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(slot, "id = ?", id).Error
	}); err != nil {
		return nil, ErrInternalServer
	}

	return slot, nil
}
