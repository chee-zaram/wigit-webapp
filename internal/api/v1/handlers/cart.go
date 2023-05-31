package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/internal/db"
	"gorm.io/gorm"
)

// CustomerPostToCart adds a new item to the database. It is equivalent to adding an item
// to a cart.
func CustomerPostToCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	user := _user.(*db.User)
	_item := new(db.Item)

	if err := ctx.ShouldBindJSON(_item); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_item.UserID = user.ID
	product, err := getProductFromDB(*_item.ProductID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Make sure quantity for item to be carted is not more than product in stock.
	if err := validateItemQuantity(_item, product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Compute amount for item
	amount := product.Price.Mul(decimal.NewFromInt(int64(*_item.Quantity)))
	_item.Amount = &amount

	if err := _item.SaveToDB(); err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Item already exists for this user"})
		return
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	item, err := getItemFromDB(*_item.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Item created successfully",
		"data": item,
	})
}

// validateItemQuantity verifies the quantity for the item to be added is valid.
func validateItemQuantity(item *db.Item, product *db.Product) error {
	if item.Quantity == nil {
		return errors.New("Item quantity must be provided")
	} else if *item.Quantity > *product.Stock {
		if *product.Stock == 0 {
			return errors.New("Cannot add to cart. Product is out of stock")
		}
		*item.Quantity = *product.Stock
	} else if *item.Quantity == 0 {
		return errors.New("Item quantity cannot be 0")
	}

	return nil
}

// getItemFromDB retrieves an item from the database based on the id.
// It returns an error if any occured.
func getItemFromDB(id string) (*db.Item, error) {
	item := new(db.Item)

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.First(item, "id = ?", id).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Item not found in database")
	} else if err != nil {
		return nil, err
	}

	return item, nil
}

// CustomerDeleteFromCart deletes an item with given id from the database.
func CustomerDeleteFromCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	user := _user.(*db.User)
	id := ctx.Param("item_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Item ID missing"})
		return
	}

	// Will delete only if the item is in the cart i.e order_id is NULL.
	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM items WHERE id = ? AND user_id = ? AND order_id is NULL`, id, *user.ID).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Item deleted successfully",
	})
}

// CustomerGetCart retrieves a slice of all items in a user's cart i.e without an order_id.
func CustomerGetCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	user := _user.(*db.User)
	var items []db.Item

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at asc").Where("user_id = ?", *user.ID).Where("order_id is NULL").Find(&items).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

// CustomerClearCart clears all items in a user's cart.
func CustomerClearCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	user := _user.(*db.User)

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM items WHERE user_id = ? AND order_id is NULL`, *user.ID).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Cart cleared successfully",
	})
}
