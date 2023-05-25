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

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

// CustomerPostBooking adds a new booking to the database for the customer.
func CustomerPostBooking(ctx *gin.Context) {
	_booking := new(models.Booking)

	if err := ctx.ShouldBindJSON(_booking); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*models.User)

	service, err := getServiceFromDB(*_booking.ServiceID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_booking.Amount = service.Price
	user.Bookings = append(user.Bookings, *_booking)
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	}); err != nil {
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
func getBookingFromDB(id string) (*models.Booking, error) {
	booking := new(models.Booking)

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Slot").Preload("Service").First(booking, "id = ?", id).Error
	}); err != nil {
		return nil, err
	}

	return booking, nil
}
