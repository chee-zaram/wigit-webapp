package db

import (
	"errors"

	"gorm.io/gorm"
)

// SaveToDB saves the current product to the database.
func (product *Product) SaveToDB() error {
	if product == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(product).Error
	}); err != nil {
		return err
	}

	return nil
}

// LoadFromDB fills up the product with the information from the database.
func (product *Product) LoadFromDB(id string) error {
	if product == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.First(product, "id = ?", id).Error
	}); err != nil {
		return err
	}

	return nil
}

// Reload saves the product to the database and loads the updated version.
func (product *Product) Reload() error {
	if product == nil {
		return ErrNilPointer
	}

	if err := product.SaveToDB(); err != nil {
		return err
	}

	if err := product.LoadFromDB(*product.ID); err != nil {
		return err
	}

	return nil
}

// AllProducts returns all the products in the database, ordered by last updated.
func AllProducts() ([]Product, error) {
	var products []Product

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").Find(&products).Error
	}); err != nil {
		return nil, err
	}

	return products, nil
}

// TrendingProducts returns the top ten trending products.
func TrendingProducts(items []Item) ([]Product, error) {
	products := make([]Product, 0)

	for _, item := range items {
		product := new(Product)
		if err := Connector.Query(func(tx *gorm.DB) error {
			return tx.First(product, "id = ?", *item.ProductID).Error
		}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		} else if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

// GetProductsByCategory gets all the products belonging to the given category.
func GetProductsByCategory(category string) ([]Product, error) {
	var products []Product

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("category = ?", category).Find(&products).Error
	}); err != nil {
		return nil, err
	}

	return products, nil
}

// DeleteProduct deletes a product from the database.
func DeleteProduct(id string) error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM products WHERE id = ?`, id).Error
	}); err != nil {
		return err
	}

	return nil
}

// GetProductsByName gets all products in the database containing the substring
// `name` using a case insensitive search.
func GetProductsByName(name string) ([]Product, error) {
	var products []Product

	if err := Connector.Query(func(tx *gorm.DB) error {
		// This insensitive search is dependent on the character set the database
		// is configured in, in this case `utf8mb4`.
		return tx.
			Where("name COLLATE utf8mb4_general_ci LIKE ?", "%"+name+"%").
			Order("name").
			Find(&products).Error
	}); err != nil {
		return nil, err
	}
	return products, nil
}
