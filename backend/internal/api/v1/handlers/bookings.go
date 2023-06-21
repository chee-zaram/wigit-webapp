package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/backend/internal/db"
)

// BookingRequest binds to the request body on post to bookings routes.
type BookingRequest struct {
	// ServiceID is the ID of the service requested.
	ServiceID *string `json:"service_id" binding:"required"`
	// SlotID is the ID of the slot the user is requesting to be booked for.
	SlotID *string `json:"slot_id" binding:"required"`
}

// validBookingStatuses is a list of all the valid status for a booking.
var validBookingStatuses = []string{
	db.Pending, db.Paid, db.Fulfilled, db.Cancelled,
}

// GetCustomerBookings Customer get all bookings
//
//	@Summary	Retrieves all their bookings from the database
//	@Tags		bookings
//	@Produce	json
//	@Param		Authorization	header		string					true	"Authorization token in the format 'Bearer <token>'"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/bookings [get]
func GetCustomerBookings(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	bookings, err := db.CustomerBookings(*loggedInUser.ID)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

// CustomerGetBooking Gets a customer's booking with given ID
//
//	@Summary	Retrieves a booking with given ID from database
//	@Tags		bookings
//	@Produce	json
//	@Param		Authorization	header		string					true	"Authorization token in the format 'Bearer <token>'"
//	@Param		booking_id		path		string					true	"ID of the booking to get"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Router		/bookings/{booking_id} [get]
func CustomerGetBooking(ctx *gin.Context) {
	bookingID := ctx.Param("booking_id")
	if bookingID == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("The booking ID parameter is missing or empty"))
		return
	}

	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	for _, userBooking := range loggedInUser.Bookings {
		if strings.HasPrefix(*userBooking.ID, bookingID) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": userBooking,
			})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": "No booking with given ID for user",
	})
}

// CustomerPostBooking adds a new booking to the database for the customer.
//
//	@Summary	Adds a new booking to the database for the authorized customer.
//	@Tags		bookings
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		booking			body		BookingRequest			true	"A new customer booking"
//	@Success	201				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/bookings [post]
func CustomerPostBooking(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	bookingRequest := new(BookingRequest)
	if err := ctx.ShouldBindJSON(bookingRequest); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	service := new(db.Service)
	if err := service.LoadFromDB(*bookingRequest.ServiceID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	booking := newBooking(bookingRequest)
	// Set teh booking amount to the price of the service
	booking.Amount = service.Price
	// Add the new booking to the customer's bookings history
	loggedInUser.Bookings = append(loggedInUser.Bookings, *booking)
	if err := loggedInUser.SaveToDB(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := booking.LoadFromDB(
		*loggedInUser.Bookings[len(loggedInUser.Bookings)-1].ID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Booking created successfully",
		"data": booking,
	})
}

// newBooking creates a new db.Booking object from the BookingRequest.
func newBooking(bookingRequest *BookingRequest) *db.Booking {
	booking := new(db.Booking)
	booking.SlotID = bookingRequest.SlotID
	booking.ServiceID = bookingRequest.ServiceID

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
	_user, exists := ctx.Get("user")
	admin, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

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

	if !isValidBookingStatus(booking, status) {
		AbortCtx(ctx, http.StatusBadRequest, errors.New(
			"The status is not valid. Likely because the service is not available at the moment",
		))
		return
	}

	adminFullName := fmt.Sprintf("%s %s", *admin.FirstName, *admin.LastName)
	if err := booking.UpdateStatus(status, adminFullName); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Booking status updated successfully",
		"data": booking,
	})
}

// isValidBookingStatus checks if the new status is for updating a booking valid.
func isValidBookingStatus(booking *db.Booking, status string) bool {
	var isValidStatus bool

	for i, stat := range validBookingStatuses {
		if i == len(validBookingStatuses)-1 && stat != status {
			isValidStatus = false
			break
		}

		if stat != status {
			continue
		}

		if status == db.Paid {
			if !*booking.Service.Available {
				return false
			}

			if !*booking.Slot.IsFree {
				return false
			}
			*booking.Slot.IsFree = false
		}

		isValidStatus = true
		break
	}

	return isValidStatus
}
