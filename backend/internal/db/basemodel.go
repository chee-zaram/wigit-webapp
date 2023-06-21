package db

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeCreate runs before inserting the row in the database table.
// It makes sure the object is valid, and the ID exists.
func (baseModel *BaseModel) BeforeCreate(_ *gorm.DB) error {
	if baseModel == nil {
		return ErrNilPointer
	}

	if baseModel.ID != nil && *baseModel.ID != "" {
		return nil
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf(
			"failed to generate UUID for BaseModel in BeforeCreate hook: %w",
			err,
		)
	}

	uid := id.String()
	baseModel.ID = &uid

	return nil
}
