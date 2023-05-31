package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db"
	"gorm.io/gorm"
)

// allowedBookingStatus is a list of all the valid status for a booking.
var allowedBookingStatus = []string{"pending", "paid", "fulfilled", "cancelled"}

// AdminGetBookings retrieves all bookings in the database.
func AdminGetBookings(ctx *gin.Context) {
	var bookings []db.Booking

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").Find(&bookings).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

// CustomerGetBookings gets a list of all the customer bookings.
func CustomerGetBookings(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*db.User)
	var bookings []db.Booking

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").Where("user_id = ?", *user.ID).Find(&bookings).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

// CustomerPostBooking adds a new booking to the database for the customer.
func CustomerPostBooking(ctx *gin.Context) {
	_booking := new(db.Booking)

	if err := ctx.ShouldBindJSON(_booking); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*db.User)

	service, err := getServiceFromDB(*_booking.ServiceID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_booking.Amount = service.Price
	user.Bookings = append(user.Bookings, *_booking)
	if err := user.SaveToDB(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	booking, err := getBookingFromDB(*user.Bookings[len(user.Bookings)-1].ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Booking created successfully",
		"data": booking,
	})
}

// getBookingFromDB retrieves a booking with id from database.
func getBookingFromDB(id string) (*db.Booking, error) {
	booking := new(db.Booking)

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Slot").Preload("Service").First(booking, "id LIKE ?", "%"+id+"%").Error
	}); err != nil {
		return nil, err
	}

	return booking, nil
}

// AdminPutBooking updates the status of a booking.
func AdminPutBooking(ctx *gin.Context) {
	id := ctx.Param("booking_id")
	status := ctx.Param("status")
	if id == "" || status == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Booking ID or Status not set"})
		return
	}

	booking, err := getBookingFromDB(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !validBookingStatus(booking, status) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "The status is not valid. Likely because the service is not available at the moment",
		})
		return
	}

	booking.Status = &status
	if err := booking.SaveToDB(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

// CustomerGetBooking retrieves booking with given id for a customer.
func CustomerGetBooking(ctx *gin.Context) {
	booking_id := ctx.Param("booking_id")
	if booking_id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Booking ID not set"})
		return
	}

	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*db.User)

	for _, booking := range user.Bookings {
		if strings.Contains(*booking.ID, booking_id) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": booking,
			})
		}
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No booking with given ID for user"})
}

// AdminGetBooking retrieves booking with given id from the database for an admin.
func AdminGetBooking(ctx *gin.Context) {
	booking_id := ctx.Param("booking_id")
	if booking_id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Booking ID not set"})
		return
	}

	booking, err := getBookingFromDB(booking_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booking,
	})
}
