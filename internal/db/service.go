package db

import "gorm.io/gorm"

// SaveToDB saves the current service to the database.
func (service *Service) SaveToDB() error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(service).Error
	}); err != nil {
		return err
	}

	return nil
}
