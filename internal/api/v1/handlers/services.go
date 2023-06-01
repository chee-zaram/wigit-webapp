package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wigit-gh/webapp/internal/db"
	"gorm.io/gorm"
)

// GetServices retrieves a list of all available services.
func GetServices(ctx *gin.Context) {
	var services []db.Service

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at DESC").Find(&services).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
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

	service := new(db.Service)
	if err := db.Connector.Query(func(tx *gorm.DB) error {
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

// AdminPostServices adds a new service to the database.
func AdminPostServices(ctx *gin.Context) {
	_service := new(db.Service)

	if err := ctx.ShouldBindJSON(_service); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateServicesData(_service); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := _service.SaveToDB(); err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Service with this name already exists"})
		return
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	service, err := getServiceFromDB(*_service.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Service created successfully",
		"data": service,
	})
}

// validateServicesData checks to make sure the data to be added is valid.
func validateServicesData(service *db.Service) error {
	if service.Price == nil || service.Price.Sign() < 0 {
		return errors.New("Price must be set and cannot be a negative value")
	}

	return nil
}

// getServiceFromDB retrieves a service from the database based on the id provided.
func getServiceFromDB(id string) (*db.Service, error) {
	service := new(db.Service)

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.First(service, "id = ?", id).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to add service to the database")
	} else if err != nil {
		return nil, ErrInternalServer
	}

	return service, nil
}

// AdminDeleteServices handles the deletion of a service from the database by an admin.
func AdminDeleteServices(ctx *gin.Context) {
	id := ctx.Param("service_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidServiceID.Error()})
		return
	}

	if err := deleteServiceFromDB(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Service deleted successfully",
	})
}

// deleteServiceFromDB sends a delete query to the database to delete service with given id.
func deleteServiceFromDB(id string) error {
	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM services WHERE id = ?`, id).Error
	}); err != nil { // Exec does not return a ErrRecordNotFound so no need to check
		return err
	}

	return nil
}

// AdminPutServices handles put requests to the services endpoint.
func AdminPutServices(ctx *gin.Context) {
	_service := new(db.Service)
	id := ctx.Param("service_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidServiceID.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(_service); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := getServiceFromDB(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := updateServiceInDB(service, _service); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Service updated successfully",
		"data": service,
	})
}

// updateServiceInDB updates the service's information in the database.
func updateServiceInDB(dbService, newService *db.Service) error {
	dbService.Name = newService.Name
	dbService.Description = newService.Description
	dbService.Price = newService.Price
	dbService.Available = newService.Available

	if err := dbService.SaveToDB(); err != nil {
		return err
	}

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.First(dbService, "id = ?", *dbService.ID).Error
	}); err != nil {
		return err
	}

	return nil
}

// GetTrendingServices retrieves a list of 5 trending services.
func GetTrendingServices(ctx *gin.Context) {
	var bookings []db.Booking

	bookings, err := sortBookingsByService()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	services, err := getTrendingServices(bookings)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": services,
	})
}

// sortBookingsByService gets service ids of the top 10 services booked in the last
// 7 days by quering the bookings table.
func sortBookingsByService() ([]db.Booking, error) {
	var bookings []db.Booking

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Table("bookings").
			Select("service_id, COUNT(*) as total_bookings").
			Where("created_at >= ?", time.Now().UTC().AddDate(0, 0, -7)).
			Group("service_id").
			Order("total_bookings DESC").
			Limit(10).
			Scan(&bookings).Error
	}); err != nil {
		return nil, err
	}

	return bookings, nil
}

// getTrendingServices retrieves the top 10 services in the last week if available.
func getTrendingServices(bookings []db.Booking) ([]db.Service, error) {
	var services []db.Service
	for _, booking := range bookings {
		service := new(db.Service)
		if err := db.Connector.Query(func(tx *gorm.DB) error {
			return tx.First(service, "id = ?", *booking.ServiceID).Error
		}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		} else if err != nil {
			return nil, err
		}
		services = append(services, *service)
	}

	return services, nil
}
