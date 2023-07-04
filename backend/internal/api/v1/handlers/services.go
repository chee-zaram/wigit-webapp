package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/chee-zaram/wigit-webapp/backend/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// ServiceRequest binds to the request body during a post or put request.
type ServiceRequest struct {
	// Name is the name of the service.
	Name *string `json:"name" binding:"required,min=3,max=45"`

	// Description is a brief description of the service.
	Description *string `json:"description" binding:"required,min=5,max=1024"`

	// Price is the cost of the service.
	Price *decimal.Decimal `json:"price" binding:"required"`

	// Available says whether the service is available or not.
	Available *bool `json:"available" binding:"required"`
}

// cleanUp removes all leading and trailing whitespace from the data.
func (s *ServiceRequest) cleanUp() {
	if s == nil {
		return
	}

	*s.Name = strings.TrimSpace(*s.Name)
	*s.Description = strings.TrimSpace(*s.Description)
}

// validateData checks to make sure the data to be added is valid.
func (service *ServiceRequest) validateData() error {
	if service.Price == nil || service.Price.Sign() < 0 {
		return errors.New("Price must be set and cannot be a negative value")
	}

	return nil
}

// GetServices	Gets all services
//
//	@Summary	Retrieves a list of all product objects
//	@Tags		services
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}	"data"
//	@Failure	500	{object}	map[string]interface{}	"error"
//	@Router		/services [get]
func GetServices(ctx *gin.Context) {
	services, err := db.AllServices()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": services,
	})
}

// GetServiceByID	Gets a service by ID
//
//	@Summary	Retrieves the service with the given ID
//	@Tags		services
//	@Produce	json
//	@Param		service_id	path		string					true	"Service ID"
//	@Success	200			{object}	map[string]interface{}	"data"
//	@Failure	400			{object}	map[string]interface{}	"error"
//	@Failure	500			{object}	map[string]interface{}	"error"
//	@Router		/services/{service_id} [get]
func GetServiceByID(ctx *gin.Context) {
	id := ctx.Param("service_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidServiceID)
		return
	}

	service := new(db.Service)
	if err := service.LoadFromDB(id); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidServiceID)
		return
	} else if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": service,
	})
}

// AdminPostService	Add service
//
//	@Summary	Allows the admin add services to the database
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		service			body		ServiceRequest			true	"Data of the service to add"
//	@Success	201				{object}	map[string]interface{}	"data, msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/services [post]
func AdminPostService(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	admin, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	_service := new(ServiceRequest)
	if err := ctx.ShouldBindJSON(_service); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	_service.cleanUp()
	if err := _service.validateData(); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	adminName := fmt.Sprintf("%s %s", *admin.FirstName, *admin.LastName)
	service := newService(_service, adminName, false)
	if err := service.SaveToDB(); err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Service with this name already exists"))
		return
	} else if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := service.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Service created successfully",
		"data": service,
	})
}

// newService fills up the fields for db.Service object.
func newService(newService *ServiceRequest, adminName string, exists bool) *db.Service {
	service := new(db.Service)
	service.Name = newService.Name
	service.Description = newService.Description
	service.Available = newService.Available
	service.Price = newService.Price
	if !exists {
		service.AddedBy = adminName
		msg := fmt.Sprintf(
			"service [%s] added by [%s]. price [%s]",
			*service.Name, adminName, service.Price,
		)
		log.Info().Msg(msg)
	} else {
		service.UpdatedBy = adminName
		msg := fmt.Sprintf(
			"service [%s] update by [%s]. price [%s]",
			*service.Name, adminName, *service.Price,
		)
		log.Info().Msg(msg)
	}

	return service
}

// AdminDeleteService	Deletes a service
//
//	@Summary	Allows admins delete a service from the database
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		service_id		path		string					true	"Service ID to delete"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/services/{service_id} [delete]
func AdminDeleteService(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	admin, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	id := ctx.Param("service_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidServiceID)
		return
	}

	if err := db.DeleteService(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}
	adminName := fmt.Sprintf("%s %s", *admin.FirstName, *admin.LastName)
	msg := fmt.Sprintf("service [%s] deleted by admin [%s]", id, adminName)
	log.Info().Msg(msg)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Service deleted successfully",
	})
}

// AdminPutService		Update product
//
//	@Summary	Allows the admin update the service with given service_id
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		service_id		path		string					true	"Service ID to update"
//	@Param		service			body		ServiceRequest			true	"Data of the service to update"
//	@Success	200				{object}	map[string]interface{}	"data, msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/services/{service_id} [put]
func AdminPutService(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	admin, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	_service := new(ServiceRequest)
	id := ctx.Param("service_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidServiceID)
		return
	}

	if err := ctx.ShouldBindJSON(_service); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	_service.cleanUp()
	if err := _service.validateData(); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	service := new(db.Service)
	if err := service.LoadFromDB(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	created_at := service.CreatedAt
	adminName := fmt.Sprintf("%s %s", *admin.FirstName, *admin.LastName)
	service = newService(_service, adminName, true)
	service.ID = &id
	service.CreatedAt = created_at

	if err := service.SaveToDB(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := service.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Service updated successfully",
		"data": service,
	})
}

// GetTrendingServices	Get the trending services
//
//	@Summary	Retrieves a list of the top ten trending services
//	@Tags		services
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}	"data"
//	@Failure	500	{object}	map[string]interface{}	"error"
//	@Router		/services/trending [get]
func GetTrendingServices(ctx *gin.Context) {
	bookings, err := db.SortBookingsByService()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	services, err := db.GetTrendingServices(bookings)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": services,
	})
}
