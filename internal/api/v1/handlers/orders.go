package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/internal/db"
	"gorm.io/gorm"
)

// allowedOrderStatus is a slice of status allowed to be set for an order.
var allowedOrderStatus = []string{"pending", "paid", "shipped", "delivered", "cancelled"}

// CustomerPostOrders adds a new order to the database for a given customer.
func CustomerPostOrders(ctx *gin.Context) {
	_order := new(db.Order)
	if err := ctx.ShouldBindJSON(_order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*db.User)

	items, err := getItemsInCart(*user.ID)
	if err != nil || items == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_order.Items = items
	_order.TotalAmount = getOrderTotal(items)

	user.Orders = append(user.Orders, *_order)
	if err := user.SaveToDB(); err != nil {
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
func getItemsInCart(user_id string) ([]db.Item, error) {
	var items []db.Item

	if err := db.Connector.Query(func(tx *gorm.DB) error {
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
func getOrderFromDB(id string) (*db.Order, error) {
	order := new(db.Order)
	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Items.Product").First(order, "id = ?", id).Error
	}); err != nil {
		return nil, err
	}

	return order, nil
}

// getOrderTotal computes the sum of all the items for a given order.
func getOrderTotal(items []db.Item) *decimal.Decimal {
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

	order, err := getOrderFromDB(order_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !validOrderStatus(order, status) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Status cannot be updated. Likely because a product is out of stock"})
		return
	}

	order.Status = &status

	if err := order.SaveToDB(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Order status updated successfully",
		"data": order,
	})
}

// validOrderStatus validates the status to which an order is about to be updated.
// returns true if the status is valid, or false otherwise.
func validOrderStatus(order *db.Order, status string) bool {
	var valid bool

	for _, stat := range allowedOrderStatus {
		if stat == status {
			if status == "paid" {
				for _, item := range order.Items {
					if *item.Quantity > *item.Product.Stock {
						return false
					}
					*item.Product.Stock = *item.Product.Stock - *item.Quantity
				}
				return true
			}
			return true
		}
	}

	return valid
}

// AdminGetOrders retrieves all the orders from the database.
func AdminGetOrders(ctx *gin.Context) {
	var orders []db.Order

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").Preload("Items").Find(&orders).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// AdminGetOrdersByStatus gets all orders with a given status.
func AdminGetOrdersByStatus(ctx *gin.Context) {
	status := ctx.Param("status")
	if status == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Status not given"})
		return
	}

	var orders []db.Order

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").Where("status = ?", status).Preload("Items").Find(&orders).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// AdminGetOrdersByID gets an order with a given id from the database.
func AdminGetOrderByID(ctx *gin.Context) {
	id := ctx.Param("order_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Order ID not provided"})
		return
	}

	order := new(db.Order)
	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Items").First(order, "id LIKE ?", "%"+id+"%").Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

// CustomerGetOrders retrieves all the user's orders from the database.
func CustomerGetOrders(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*db.User)
	var orders []db.Order

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").Where("user_id = ?", *user.ID).Preload("Items").Find(&orders).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// CustomerGetOrdersByStatus retrieves all the orders for the user by their status.
func CustomerGetOrdersByStatus(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*db.User)

	status := ctx.Param("status")
	if status == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Status not given"})
		return
	}

	var orders []db.Order

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("updated_at desc").Where("user_id = ?", *user.ID).Where("status = ?", status).
			Preload("Items").Find(&orders).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// CustomerGetOrderByID retrieves a specified order for a user.
func CustomerGetOrderByID(ctx *gin.Context) {
	id := ctx.Param("order_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Order ID not provided"})
		return
	}

	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}

	user := _user.(*db.User)
	order := new(db.Order)

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("user_id = ?", *user.ID).Preload("Items").First(order, "id LIKE ?", "%"+id+"%").Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}
