package db

import "gorm.io/gorm"

// SaveToDB saves the current user to the database.
func (user *User) SaveToDB() error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	}); err != nil {
		return err
	}

	return nil
}
