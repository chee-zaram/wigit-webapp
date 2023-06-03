package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db"
)

// NewSlot binds to the request json body during a post to /slots
type NewSlot struct {
	// DateString is the date as a string. Format `Wednesday, 06 Jan 1999`
	DateString *string `json:"date_string" binding:"required"`
	// TimeString is the time as a string. Format `04:00 AM`
	TimeString *string `json:"time_string" binding:"required"`
	// IsFree is a boolean that says if the slot is free or not.
	IsFree *bool `json:"is_free" binding:"required"`
}

// GetSlots Gets a list of all available slots
//
//	@Summary	Retrieves a list of all slot objects
//	@Tags		slots
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}	"data"
//	@Failure	500	{object}	map[string]interface{}	"error"
//	@Router		/slots [get]
func GetSlots(ctx *gin.Context) {
	slots, err := db.AllSlots()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": slots,
	})
}

// AdminPostSlots Add a new slot
//
//	@Summary	Allows the admin add slots to the database
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		product			body		NewSlot					true	"Add product"
//	@Success	201				{object}	map[string]interface{}	"data, msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/slots [post]
func AdminPostSlots(ctx *gin.Context) {
	_slot := new(NewSlot)

	if err := ctx.ShouldBindJSON(_slot); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	slot := newSlot(_slot)
	if err := slot.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Slot created successfully",
		"data": slot,
	})
}

// newSlot fills a new db.Slot object with fields from NewSlot.
func newSlot(newSlot *NewSlot) *db.Slot {
	slot := new(db.Slot)
	slot.DateString = newSlot.DateString
	slot.TimeString = newSlot.TimeString
	slot.IsFree = newSlot.IsFree

	return slot
}

// AdminDeleteSlots Deletes a slot
//
//	@Summary	Allows admins delete a slot from the database
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		slot_id			path		string					true	"Slot ID to delete"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/slots/{slot_id} [delete]
func AdminDeleteSlots(ctx *gin.Context) {
	id := ctx.Param("slot_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidSlotID)
		return
	}

	if err := db.DeleteSlot(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Slot deleted successfully",
	})
}
