package db

import "gorm.io/gorm"

// SaveToDB saves the current booking to the database.
func (booking *Booking) SaveToDB() error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(booking).Error
	}); err != nil {
		return err
	}

	return nil
}
