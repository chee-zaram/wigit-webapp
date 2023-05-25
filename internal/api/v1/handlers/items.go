package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

func PostItems(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	user := _user.(*models.User)
	_item := new(models.Item)

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

	if *_item.Quantity > *product.Stock {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Item quantity more than product stock"})
		return
	}
	// Compute amount for item
	amount := product.Price.Mul(decimal.NewFromInt(int64(*_item.Quantity)))
	_item.Amount = &amount

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Create(_item).Error
	}); err != nil && strings.Contains(err.Error(), "Duplicate entry") {
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

// getItemFromDB retrieves an item from the database based on the id.
// It returns an error if any occured.
func getItemFromDB(id string) (*models.Item, error) {
	item := new(models.Item)

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(item, "id = ?", id).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Item not found in database")
	} else if err != nil {
		return nil, err
	}

	return item, nil
}
