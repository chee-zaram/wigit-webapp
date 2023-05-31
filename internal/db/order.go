package db

import "gorm.io/gorm"

// SaveToDB saves the current order to the database.
func (order *Order) SaveToDB() error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(order).Error
	}); err != nil {
		return err
	}

	return nil
}
