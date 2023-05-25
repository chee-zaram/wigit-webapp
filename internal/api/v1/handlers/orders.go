package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// allowedOrderStatus is a slice of status allowed to be set for an order.
var allowedOrderStatus = []string{"pending", "paid", "shipped", "delivered", "cancelled"}

// PostOrders adds a new order to the database.
func PostOrders(ctx *gin.Context) {
	_order := new(models.Order)
	if err := ctx.ShouldBindJSON(_order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*models.User)

	items, err := getItemsInCart(*user.ID)
	if err != nil || items == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_order.Items = items
	_order.TotalAmount = getOrderTotal(items)

	user.Orders = append(user.Orders, *_order)
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order, err := getOrderFromDB(*user.Orders[len(user.Orders)-1].ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Order created successfully",
		"data": order,
	})
}

// getItemsInCart returns all the items in a user's cart.
func getItemsInCart(user_id string) ([]models.Item, error) {
	var items []models.Item

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Where("user_id = ?", user_id).Where("order_id is NULL").Find(&items).Error
	}); err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("No items in cart")
	}

	for _, item := range items {
		product, err := getProductFromDB(*item.ProductID)
		if err != nil {
			return nil, err
		}

		if *item.Quantity > *product.Stock {
			*item.Quantity = *product.Stock
		}

		// Compute amount for item
		amount := product.Price.Mul(decimal.NewFromInt(int64(*item.Quantity)))
		item.Amount = &amount
	}

	return items, nil
}

// getOrderFromDB retrieves an order from the database with given id.
func getOrderFromDB(id string) (*models.Order, error) {
	order := new(models.Order)
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Items").First(order, "id = ?", id).Error
	}); err != nil {
		return nil, err
	}

	return order, nil
}

// getOrderTotal computes the sum of all the items for a given order.
func getOrderTotal(items []models.Item) *decimal.Decimal {
	var totalAmount decimal.Decimal
	for _, item := range items {
		totalAmount = totalAmount.Add(*item.Amount)
	}

	return &totalAmount
}

// AdminPutOrders updates the status of an order.
func AdminPutOrders(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	if order_id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No order ID"})
		return
	}

	status := ctx.Param("status")
	if status == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No status specified"})
		return
	}

	if !validStatus(status) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Status is not valid"})
		return
	}

	order, err := getOrderFromDB(order_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order.Status = &status

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Save(order).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Order status updated successfully",
		"data": order,
	})
}

// validStatus validates the status to which an order is about to be updated.
// returns true if the status is valid, or false otherwise.
func validStatus(status string) bool {
	var valid bool

	for _, stat := range allowedOrderStatus {
		if stat == status {
			return true
		}
	}

	return valid
}
