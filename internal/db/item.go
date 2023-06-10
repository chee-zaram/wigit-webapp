package db

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// SaveToDB saves the current item to the database.
func (item *Item) SaveToDB() error {
	if item == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(item).Error
	}); err != nil {
		return err
	}

	return nil
}

// LoadFromDB loads the database version of the item into the current object.
func (item *Item) LoadFromDB(id string) error {
	if item == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Product").Where("order_id IS NULL").First(item, "id = ?", id).Error
	}); err != nil {
		return err
	}

	return nil
}

// Reload saves the item to the database and loads the updated version.
func (item *Item) Reload() error {
	if item == nil {
		return ErrNilPointer
	}

	if err := item.SaveToDB(); err != nil {
		return err
	}

	if err := item.LoadFromDB(*item.ID); err != nil {
		return err
	}

	return nil
}

// IsDuplicate says if the item to be added to cart is a duplicate.
func (item *Item) IsDuplicate() bool {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("product_id = ?", *item.ProductID).
			Where("user_id = ?", *item.UserID).
			Where("order_id IS NULL").
			First(&Item{}).Error
	}); err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

// TrendingItems returns top ten trending products in the last 7 days by ids.
func TrendingItems() ([]Item, error) {
	var items []Item

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Table("items").Select("product_id, SUM(quantity) as total_orders").
			Where("created_at >= ?", time.Now().UTC().AddDate(0, 0, -7)).
			Group("product_id").
			Order("total_orders DESC").
			Limit(10).
			Scan(&items).Error
	}); err != nil {
		return nil, err
	}

	return items, nil
}

// GetItemsInCart gets items in a user's cart for viewing.
func GetItemsInCart(userID string) ([]Item, error) {
	var items []Item

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at asc").
			Where("user_id = ?", userID).
			Where("order_id is NULL").
			Preload("Product").
			Find(&items).Error
	}); err != nil {
		return nil, err
	}

	return items, nil
}

// GetItemsInCartForOrder returns all the items in a user's cart when they post an order.
func GetItemsInCartForOrder(userID string) ([]Item, error) {
	items, err := GetItemsInCart(userID)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("No items in cart")
	}

	if err := validateStock(items); err != nil {
		return nil, err
	}

	return items, nil
}

// validateStock makes sure every item in cart available and in sufficient quantity.
func validateStock(items []Item) error {
	for _, item := range items {
		product := new(Product)
		if err := product.LoadFromDB(*item.ProductID); err != nil {
			return err
		}

		if *product.Stock == 0 {
			return errors.New("Product is out of stock")
		}

		if *item.Quantity > *product.Stock {
			*item.Quantity = *product.Stock
		}

		// Compute amount for item
		amount := product.Price.Mul(decimal.NewFromInt(*item.Quantity))
		item.Amount = &amount
	}

	return nil
}

// DeleteItem deletes the item with given `id` from the database.
func DeleteItem(userID, id string) error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM items WHERE id = ? AND user_id = ? AND order_id is NULL`, id, userID).Error
	}); err != nil {
		return err
	}

	return nil
}

// ClearCart deletes all items in the given user's cart.
func ClearCart(userID string) error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM items WHERE user_id = ? AND order_id is NULL`, userID).Error
	}); err != nil {
		return err
	}

	return nil
}
