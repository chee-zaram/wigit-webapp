package db

import (
	"errors"

	"gorm.io/gorm"
)

// SaveToDB saves the current service to the database.
func (service *Service) SaveToDB() error {
	if service == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(service).Error
	}); err != nil {
		return err
	}

	return nil
}

// LoadFromDB fills up the service with the information from the database.
func (service *Service) LoadFromDB(id string) error {
	if service == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.First(service, "id = ?", id).Error
	}); err != nil {
		return err
	}

	return nil
}

// Reload reloads the current service from the database with its ID.
func (service *Service) Reload() error {
	if service == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.First(service, "id = ?", *service.ID).Error
	}); err != nil {
		return err
	}

	return nil
}

// AllServices retturns all services currently in the database, ordered by last updated.
func AllServices() ([]Service, error) {
	var services []Service

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at DESC").Find(&services).Error
	}); err != nil {
		return nil, err
	}

	return services, nil
}

// DeleteService deletes a service with ID from database.
func DeleteService(id string) error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM services WHERE id = ?`, id).Error
	}); err != nil {
		return err
	}

	return nil
}

// GetTrendingServices retrieves the top 10 services in the last week if available.
func GetTrendingServices(bookings []Booking) ([]Service, error) {
	var services []Service
	for _, booking := range bookings {
		service := new(Service)
		if err := Connector.Query(func(tx *gorm.DB) error {
			return tx.First(service, "id = ?", *booking.ServiceID).Error
		}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		} else if err != nil {
			return nil, err
		}
		services = append(services, *service)
	}

	return services, nil
}
