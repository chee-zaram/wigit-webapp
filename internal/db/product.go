package db

import "gorm.io/gorm"

// SaveToDB saves the current product to the database.
func (product *Product) SaveToDB() error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(product).Error
	}); err != nil {
		return err
	}

	return nil
}
