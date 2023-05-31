package db

import "gorm.io/gorm"

// SaveToDB saves the current slot to the database.
func (slot *Slot) SaveToDB() error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(slot).Error
	}); err != nil {
		return err
	}

	return nil
}
