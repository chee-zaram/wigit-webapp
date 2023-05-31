package migration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wigit-gh/webapp/internal/db/models"
)

// TestAutoMigrate runs tests for the AutoMigrate function.
func TestGetSchemas(t *testing.T) {
	assert := assert.New(t)

	user, order, booking, slot, item, product, service := GetSchemas()
	assert.IsType(&models.User{}, user)
	assert.IsType(&models.Order{}, order)
	assert.IsType(&models.Booking{}, booking)
	assert.IsType(&models.Slot{}, slot)
	assert.IsType(&models.Item{}, item)
	assert.IsType(&models.Product{}, product)
	assert.IsType(&models.Service{}, service)
}
