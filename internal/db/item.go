package db

import "gorm.io/gorm"

// SaveToDB saves the current item to the database.
func (item *Item) SaveToDB() error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(item).Error
	}); err != nil {
		return err
	}

	return nil
}
