package db

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeCreate runs before inserting the row in the database table.
// It makes sure the object is valid, and the ID exists.
func (basemodel *BaseModel) BeforeCreate(_ *gorm.DB) error {
	if basemodel == nil {
		return fmt.Errorf("failed to run BeforeCreate hook for BaseModel")
	}

	if basemodel.ID != nil && *basemodel.ID != "" {
		return nil
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to run BeforeCreate hook for BaseModel: %w", err)
	}

	uid := id.String()
	basemodel.ID = &uid

	return nil
}
