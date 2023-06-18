package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/backend/internal/db"
)

// NewOrder binds to the json body during a post request for a new order.
type NewOrder struct {
	// DeliveryMethod is the method in which the order should be deliver.
	//
	// Allowed values are `pickup`, `delivery`.
	DeliveryMethod *string `json:"delivery_method" binding:"required,max=45"`
	// ShippingAddress is the address in which the current order should be delivered
	// if the DeliveryMethod is `delivery`.
	ShippingAddress *string `json:"shipping_address"`
}

// validate validates the values of the fields in the NewOrder body.
func (order *NewOrder) validate() error {
	if order == nil {
		return db.ErrNilPointer
	}

	if *order.DeliveryMethod != "pickup" && *order.DeliveryMethod != "delivery" {
		return errors.New("Invalid delivery option")
	}

	return nil
}

// allowedOrderStatus is a slice of status allowed to be set for an order.
var allowedOrderStatus = []string{
	db.Pending, db.Paid, db.Shipped, db.Delivered, db.Cancelled,
}

// CustomerGetOrders Get all given customer's orders.
//
//	@Summary	Allows the customer retrieve all their orders from the database.
//	@Tags		orders
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/orders [get]
func CustomerGetOrders(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	orders, err := db.CustomerOrders(*user.ID)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// CustomerGetOrdersByStatus Get orders by status.
//
//	@Summary	Allows the customer retrieve all the orders from the database with given status.
//	@Tags		orders
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		status			path		string					true	"The status of orders to retrieve"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/orders/status/{status} [get]
func CustomerGetOrdersByStatus(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	status := ctx.Param("status")
	if status == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrStatusCtx)
		return
	}

	orders, err := db.CustomerOrdersByStatus(*user.ID, status)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// CustomerGetOrder Customer get an order with given ID
//
//	@Summary	Allows customer retrieve an order with given ID from database
//	@Tags		orders
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		order_id		path		string					true	"ID of the order to get"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/orders/{order_id} [get]
func CustomerGetOrder(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}
	id := ctx.Param("order_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Order ID not provided"))
		return
	}

	order := new(db.Order)
	if err := order.CustomerLoadFromDB(*user.ID, id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

// CustomerPostOrders Add order to the database
//
//	@Summary	Allows the current user place a new order
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		order			body		NewOrder				true	"A new customer order"
//	@Success	201				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/orders [post]
func CustomerPostOrders(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	_order := new(NewOrder)
	if err := ctx.ShouldBindJSON(_order); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	if err := _order.validate(); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	items, err := db.GetItemsInCartForOrder(*user.ID)
	if err != nil || items == nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	order := newOrder(_order, items)
	if *order.DeliveryMethod == "delivery" && order.ShippingAddress == nil {
		order.ShippingAddress = user.Address
	}
	user.Orders = append(user.Orders, *order)

	if err := user.SaveToDB(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := order.LoadFromDB(*user.Orders[len(user.Orders)-1].ID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Order created successfully",
		"data": order,
	})
}

// newOrder fills up fields in the db.Order object from NewOrder and returns it.
func newOrder(newOrder *NewOrder, items []db.Item) *db.Order {
	order := new(db.Order)
	order.DeliveryMethod = newOrder.DeliveryMethod
	order.Items = items
	order.TotalAmount = getOrderTotal(items)

	if *newOrder.DeliveryMethod == "delivery" {
		if newOrder.ShippingAddress != nil && *newOrder.ShippingAddress != "" {
			order.ShippingAddress = newOrder.ShippingAddress
		}
	}

	return order
}

// getOrderTotal computes the sum of all the items for a given order.
func getOrderTotal(items []db.Item) *decimal.Decimal {
	var totalAmount decimal.Decimal
	for _, item := range items {
		totalAmount = totalAmount.Add(*item.Amount)
	}

	return &totalAmount
}

// AdminGetOrders Get all database orders.
//
//	@Summary	Allows admin retrieves all orders from the database
//	@Tags		admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/orders [get]
func AdminGetOrders(ctx *gin.Context) {
	orders, err := db.AllOrders()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// AdminGetOrdersByStatus Get all orders with given status
//
//	@Summary	Allows admin retrieves all orders with given status from the database
//	@Tags		admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		status			path		string					true	"Status of orders to retrieve"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/orders/status/{status} [get]
func AdminGetOrdersByStatus(ctx *gin.Context) {
	status := ctx.Param("status")
	if status == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("No status specified"))
		return
	}

	orders, err := db.OrdersByStatus(status)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

// AdminGetOrder Get order with ID
//
//	@Summary	Allows admin retrieve an order with given ID from database
//	@Tags		admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		order_id		path		string					true	"ID of the order to get"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/orders/{order_id} [get]
func AdminGetOrder(ctx *gin.Context) {
	id := ctx.Param("order_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Order ID not provided"))
		return
	}

	order := new(db.Order)
	if err := order.LoadFromDB(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

// AdminPutOrders Update the status of an order.
//
//	@Summary		Allows admin update the status of an existing order.
//	@Description	Allowed status are `pending`(default), `paid`, `shipped`, `delivered`, `cancelled`
//	@Tags			admin
//	@Produce		json
//	@Param			Authorization	header		string					true	"Bearer <token>"
//	@Param			order_id		path		string					true	"ID of order to update"
//	@Param			status			path		string					true	"New status of the order"
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	map[string]interface{}	"error"
//	@Failure		500				{object}	map[string]interface{}	"error"
//	@Router			/admin/orders/{order_id}/{status} [put]
func AdminPutOrders(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	if order_id == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("No order ID"))
		return
	}

	status := ctx.Param("status")
	if status == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("No status specified"))
		return
	}

	order := new(db.Order)
	if err := order.LoadFromDB(order_id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	if !orderStatusIsValid(order, status) {
		AbortCtx(ctx, http.StatusBadRequest, errors.New(
			"Status cannot be updated. Likely because a product is out of stock",
		))
		return
	}

	if err := order.UpdateStatus(status); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Order status updated successfully",
		"data": order,
	})
}

// orderStatusIsValid validates the status to which an order is about to be updated.
// returns true if the status is valid, or false otherwise.
func orderStatusIsValid(order *db.Order, status string) bool {
	var valid bool

	for _, stat := range allowedOrderStatus {
		if stat == status {
			if status == db.Paid {
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
