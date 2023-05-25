package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// GetServices retrieves a list of all available services.
func GetServices(ctx *gin.Context) {
	var services []models.Service

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Find(&services).Error
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

	service := new(models.Service)
	if err := DBConnector.Query(func(tx *gorm.DB) error {
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
	_service := new(models.Service)

	if err := ctx.ShouldBindJSON(_service); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateServicesData(_service); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Create(_service).Error
	}); err != nil && strings.Contains(err.Error(), "Duplicate entry") {
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
func validateServicesData(service *models.Service) error {
	if service.Price == nil || service.Price.Sign() < 0 {
		return errors.New("Price must be set and cannot be a negative value")
	}

	return nil
}

// getServiceFromDB retrieves a service from the database based on the id provided.
func getServiceFromDB(id string) (*models.Service, error) {
	service := new(models.Service)

	if err := DBConnector.Query(func(tx *gorm.DB) error {
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
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM services WHERE id = ?`, id).Error
	}); err != nil { // Exec does not return a ErrRecordNotFound so no need to check
		return err
	}

	return nil
}

// AdminPutServices handles put requests to the services endpoint.
func AdminPutServices(ctx *gin.Context) {
	_service := new(models.Service)
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
func updateServiceInDB(dbService, newService *models.Service) error {
	dbService.Name = newService.Name
	dbService.Description = newService.Description
	dbService.Price = newService.Price
	dbService.Available = newService.Available

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Save(dbService).Error
	}); err != nil {
		return err
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(dbService, "id = ?", *dbService.ID).Error
	}); err != nil {
		return err
	}

	return nil
}
