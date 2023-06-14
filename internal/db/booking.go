package db

import (
	"time"

	"gorm.io/gorm"
)

// SaveToDB saves the current booking to the database.
func (booking *Booking) SaveToDB() error {
	if booking == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(booking).Error
	}); err != nil {
		return err
	}

	return nil
}

// LoadFromDB loads up the booking with given id into the booking object pointed
// to by `booking`.
func (booking *Booking) LoadFromDB(id string) error {
	if booking == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Slot").Preload("Service").First(booking, "id LIKE ?", "%"+id+"%").Error
	}); err != nil {
		return err
	}

	return nil
}

// UpdateStatus updates the status of an existing user booking.
func (booking *Booking) UpdateStatus(newStatus string) error {
	if booking == nil {
		return ErrNilPointer
	}

	booking.Status = &newStatus
	if err := booking.Reload(); err != nil {
		return err
	}

	return nil
}

// Reload saves and then loads up the current version of the current booking from the database.
func (booking *Booking) Reload() error {
	if booking == nil {
		return ErrNilPointer
	}

	if err := booking.SaveToDB(); err != nil {
		return err
	}

	if err := booking.LoadFromDB(*booking.ID); err != nil {
		return err
	}

	return nil
}

// AllBookings gets all the bookings currently in the database, ordered by last updated.
func AllBookings() ([]Booking, error) {
	var bookings []Booking

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").
			Preload("Slot").
			Preload("Service").
			Find(&bookings).Error
	}); err != nil {
		return nil, err
	}

	return bookings, nil
}

// CustomerBookings gets all the bookings belonging to a customer with userID
// from the database.
func CustomerBookings(userID string) ([]Booking, error) {
	var bookings []Booking

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").
			Where("user_id = ?", userID).
			Preload("Slot").
			Preload("Service").
			Find(&bookings).Error
	}); err != nil {
		return nil, err
	}

	return bookings, nil
}

// SortBookingsByService gets service ids of the top 10 services booked in the last
// 7 days by quering the bookings table.
func SortBookingsByService() ([]Booking, error) {
	var bookings []Booking

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Table("bookings").
			Select("service_id, COUNT(*) as total_bookings").
			Where("created_at >= ?", time.Now().UTC().AddDate(0, 0, -7)).
			Group("service_id").
			Order("total_bookings DESC").
			Limit(10).
			Scan(&bookings).Error
	}); err != nil {
		return nil, err
	}

	return bookings, nil
}
