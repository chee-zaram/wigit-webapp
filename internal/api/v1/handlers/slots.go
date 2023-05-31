package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// GetSlots retrieves a list of all available slots.
func GetSlots(ctx *gin.Context) {
	var slots []models.Slot

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("is_free = ?", true).Find(&slots).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": slots,
	})
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

	if err := db.Connector.Query(func(tx *gorm.DB) error {
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

// validatePostSlotsData checks that the required fields are provided.
func validatePostSlotsData(slot *models.Slot) error {
	if slot.DateString == nil || *slot.DateString == "" {
		return errors.New("DateString must be provided")
	}

	if slot.TimeString == nil || *slot.TimeString == "" {
		return errors.New("TimeString must be provided")
	}

	return nil
}

// getSlotFromDB retrieves a slot with id from the database.
func getSlotFromDB(id string) (*models.Slot, error) {
	slot := new(models.Slot)

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.First(slot, "id = ?", id).Error
	}); err != nil {
		return nil, ErrInternalServer
	}

	return slot, nil
}

// AdminDeleteSlots handles deletion of a slot by an admin.
func AdminDeleteSlots(ctx *gin.Context) {
	id := ctx.Param("slot_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidSlotID.Error()})
		return
	}

	if err := deleteSlotFromDB(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Slot deleted successfully",
	})
}

// deleteSlotFromDB deletes a slot with id from database.
func deleteSlotFromDB(id string) error {
	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM slots WHERE id = ?`, id).Error
	}); err != nil {
		return err
	}

	return nil
}
