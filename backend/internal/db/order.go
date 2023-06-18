package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// SaveToDB saves the current order to the database.
func (order *Order) SaveToDB() error {
	if order == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(order).Error
	}); err != nil {
		return err
	}

	return nil
}

// LoadFromDB loads the order with id from the database into the current order object.
func (order *Order) LoadFromDB(id string) error {
	if order == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Items.Product").
			First(order, "id LIKE ?", "%"+id+"%").Error
	}); err != nil {
		return err
	}

	return nil
}

// CustomerLoadFromDB loads the customer's order with given `id` into the object.
func (order *Order) CustomerLoadFromDB(userID, id string) error {
	if order == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("user_id = ?", userID).
			Preload("Items.Product").
			First(order, "id LIKE ?", "%"+id+"%").Error
	}); err != nil {
		return err
	}

	return nil
}

// UpdateStatus updates the status of the current order specifying the name of
// the admin responsible for the last update.
func (order *Order) UpdateStatus(status, adminName string) error {
	if order == nil {
		return ErrNilPointer
	}

	order.Status = &status
	order.UpdatedBy = adminName

	switch status {
	case Paid:
		order.PaidUpdatedBy = adminName
	case Shipped:
		order.ShippedUpdatedBy = adminName
	case Delivered:
		order.DeliveredUpdatedBy = adminName
	}

	if err := order.Reload(); err != nil {
		return err
	}

	msg := fmt.Sprintf(
		"status of order with id = [%s] updated to [%s] by [%s]",
		*order.ID, status, adminName,
	)
	log.Info().Msg(msg)

	return nil
}

// Reload saves the current order to the database and loads it back up.
func (order *Order) Reload() error {
	if err := order.SaveToDB(); err != nil {
		return err
	}

	if err := order.LoadFromDB(*order.ID); err != nil {
		return err
	}

	return nil
}

// CustomerOrders gets all orders for the current user from the database.
func CustomerOrders(userID string) ([]Order, error) {
	var orders []Order

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").
			Where("user_id = ?", userID).
			Preload("Items.Product").
			Find(&orders).Error
	}); err != nil {
		return nil, err
	}

	return orders, nil
}

// CustomerOrdersByStatus gets all orders for the customer matching the given status,
// ordered by last updated.
func CustomerOrdersByStatus(userID, status string) ([]Order, error) {
	var orders []Order

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").
			Where("user_id = ?", userID).
			Where("status = ?", status).
			Preload("Items.Product").
			Find(&orders).Error
	}); err != nil {
		return nil, err
	}

	return orders, nil
}

// AllOrders gets all order entries from the database.
func AllOrders() ([]Order, error) {
	var orders []Order

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").
			Preload("Items.Product").
			Find(&orders).Error
	}); err != nil {
		return nil, err
	}

	return orders, nil
}

// OrdersByStatus retrieves all database order entries with given status.
func OrdersByStatus(status string) ([]Order, error) {
	var orders []Order

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").
			Where("status = ?", status).
			Preload("Items.Product").
			Find(&orders).Error
	}); err != nil {
		return nil, err
	}

	return orders, nil
}
