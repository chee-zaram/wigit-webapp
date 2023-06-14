package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db"
)

// NewBooking binds to the request body on post to bookings routes.
type NewBooking struct {
	// ServiceID is the ID of the service requested.
	ServiceID *string `json:"service_id" binding:"required"`
	// SlotID is the ID of the slot the user is requesting to be booked for.
	SlotID *string `json:"slot_id" binding:"required"`
}

// allowedBookingStatus is a list of all the valid status for a booking.
var allowedBookingStatus = []string{"pending", "paid", "fulfilled", "cancelled"}

// CustomerGetBookings Customer get all bookings
//
//	@Summary	Allows customer retrieves all their bookings from the database
//	@Tags		bookings
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/bookings [get]
func CustomerGetBookings(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	bookings, err := db.CustomerBookings(*user.ID)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

// CustomerGetBooking Customer get a booking with given ID
//
//	@Summary	Allows customer retrieve a booking with given ID from database
//	@Tags		bookings
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		booking_id		path		string					true	"ID of the booking to get"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Router		/bookings/{booking_id} [get]
func CustomerGetBooking(ctx *gin.Context) {
	booking_id := ctx.Param("booking_id")
	if booking_id == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Booking ID not set"))
		return
	}

	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	for _, booking := range user.Bookings {
		if strings.Contains(*booking.ID, booking_id) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": booking,
			})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No booking with given ID for user"})
}

// CustomerPostBooking adds a new booking to the database for the customer.
//
//	@Summary	Allows customer add a new booking.
//	@Tags		bookings
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		booking			body		NewBooking				true	"A new customer booking"
//	@Success	201				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/bookings [post]
func CustomerPostBooking(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	_booking := new(NewBooking)
	if err := ctx.ShouldBindJSON(_booking); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	service := new(db.Service)
	if err := service.LoadFromDB(*_booking.ServiceID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	booking := newBooking(_booking)
	booking.Amount = service.Price
	user.Bookings = append(user.Bookings, *booking)
	if err := user.SaveToDB(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := booking.LoadFromDB(*user.Bookings[len(user.Bookings)-1].ID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Booking created successfully",
		"data": booking,
	})
}

// newBooking fills up the neccessary fields in db.Booking object from NewBooking
// object and returns it.
func newBooking(newBooking *NewBooking) *db.Booking {
	booking := new(db.Booking)
	booking.SlotID = newBooking.SlotID
	booking.ServiceID = newBooking.ServiceID

	return booking
}

// AdminGetBookings	Get all database bookings
//
//	@Summary	Allows admin retrieve all bookings from the database
//	@Tags		admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/bookings [get]
func AdminGetBookings(ctx *gin.Context) {
	bookings, err := db.AllBookings()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

// AdminGetBooking Admin get a booking with given ID
//
//	@Summary	Allows admin retrieve a booking with given ID from database
//	@Tags		admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		booking_id		path		string					true	"ID of the booking to get"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/bookings/{booking_id} [get]
func AdminGetBooking(ctx *gin.Context) {
	id := ctx.Param("booking_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Booking ID not set"))
		return
	}

	booking := new(db.Booking)
	if err := booking.LoadFromDB(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booking,
	})
}

// AdminPutBooking Update the status of a booking.
//
//	@Summary		Allows admin update the status of an existing booking.
//	@Description	Allowed status are `pending`(default), `paid`, `fulfilled`, `cancelled`
//	@Tags			admin
//	@Produce		json
//	@Param			Authorization	header		string					true	"Bearer <token>"
//	@Param			booking_id		path		string					true	"ID of booking to update"
//	@Param			status			path		string					true	"New status of the booking"
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	map[string]interface{}	"error"
//	@Failure		500				{object}	map[string]interface{}	"error"
//	@Router			/admin/bookings/{booking_id}/{status} [put]
func AdminPutBooking(ctx *gin.Context) {
	id := ctx.Param("booking_id")
	status := ctx.Param("status")
	if id == "" || status == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Booking ID or Status not set"))
		return
	}

	booking := new(db.Booking)
	if err := booking.LoadFromDB(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	if !validBookingStatus(booking, status) {
		AbortCtx(ctx, http.StatusBadRequest, errors.New(
			"The status is not valid. Likely because the service is not available at the moment",
		))
		return
	}

	if err := booking.UpdateStatus(status); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Booking status updated successfully",
		"data": booking,
	})
}

// validBookingStatus checks if the new status is for updating a booking valid.
func validBookingStatus(booking *db.Booking, status string) bool {
	var valid bool

	for _, stat := range allowedBookingStatus {
		if stat == status {
			if status == "paid" {
				if !*booking.Service.Available {
					return false
				}

				if !*booking.Slot.IsFree {
					return false
				} else {
					*booking.Slot.IsFree = false
				}
				return true
			}
			return true
		}
	}

	return valid
}
