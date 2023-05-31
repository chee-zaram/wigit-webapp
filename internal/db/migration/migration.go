package migration

import "github.com/wigit-gh/webapp/internal/db/models"

// AutoMigrate returns all the models for which schemas are to be created.
func AutoMigrate() (*models.User, *models.Order, *models.Booking, *models.Slot, *models.Item, *models.Product, *models.Service) {
	return &models.User{}, &models.Order{}, &models.Booking{}, &models.Slot{}, &models.Item{}, &models.Product{}, &models.Service{}
}
