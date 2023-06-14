package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAutoMigrate runs tests for the AutoMigrate function.
func TestGetSchemas(t *testing.T) {
	assert := assert.New(t)

	user, order, booking, slot, item, product, service := GetSchemas()
	assert.IsType(&User{}, user)
	assert.IsType(&Order{}, order)
	assert.IsType(&Booking{}, booking)
	assert.IsType(&Slot{}, slot)
	assert.IsType(&Item{}, item)
	assert.IsType(&Product{}, product)
	assert.IsType(&Service{}, service)
}
