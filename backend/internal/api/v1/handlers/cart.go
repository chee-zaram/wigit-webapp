package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/backend/internal/db"
)

// NewItem binds to the json body on post request to cart.
type NewItem struct {
	// ProductID is the id of the product the item represents.
	ProductID *string `json:"product_id" binding:"required"`
	// Quantity is the number of the product for this order.
	Quantity *int64 `json:"quantity" binding:"required"`
}

// CustomerGetCart Get all items in a user's cart
//
//	@Summary	Allows the customer retrieve all the items in their cart
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart [get]
func CustomerGetCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	items, err := db.GetItemsInCart(*user.ID)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

// CustomerPostToCart Add a new item to cart
//
//	@Summary	Allows the current user to add a new item to the cart
//	@Tags		cart
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		item			body		NewItem					true	"A new item to add to cart"
//	@Success	201				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart [post]
func CustomerPostToCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	_newItem := new(NewItem)
	if err := ctx.ShouldBindJSON(_newItem); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	item := newItem(_newItem)
	item.UserID = user.ID
	product := new(db.Product)
	if err := product.LoadFromDB(*item.ProductID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	// Make sure quantity for item to be carted is not more than product in stock.
	if err := validateItemQuantity(item, product); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	item.Amount = getItemAmount(product, *item.Quantity)

	if item.IsDuplicate() {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Item already in cart"))
		return
	}

	if err := item.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Item created successfully",
		"data": item,
	})
}

// newItem fills up a new db.Item object with bound fields from NewItem object.
func newItem(newItem *NewItem) *db.Item {
	item := new(db.Item)
	item.Quantity = newItem.Quantity
	item.ProductID = newItem.ProductID

	return item
}

// validateItemQuantity verifies the quantity for the item to be added is valid.
func validateItemQuantity(item *db.Item, product *db.Product) error {
	if item.Quantity == nil {
		return errors.New("Item quantity must be provided")
	}

	if *item.Quantity > *product.Stock {
		if *product.Stock == 0 {
			return errors.New("Cannot add to cart. Product is out of stock")
		}
		*item.Quantity = *product.Stock
	} else if *item.Quantity < 1 {
		return errors.New("Item quantity cannot be less than 1")
	}

	return nil
}

// getItemAmount Computes amount for an item.
func getItemAmount(product *db.Product, quantity int64) *decimal.Decimal {
	amount := product.Price.Mul(decimal.NewFromInt(quantity))
	return &amount
}

// CustomerPutQuantity Changes the quantity of an item in the cart
//
//	@Summary	Allows a customer edit the quantity of an item in the cart
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		item_id			path		string					true	"ID of the item"
//	@Param		quantity		path		string					true	"New item quantity"
//	@Success	200				{object}	map[string]interface{}	"msg,data"
//	@Success	400				{object}	map[string]interface{}	"error"
//	@Success	500				{object}	map[string]interface{}	"error"
//	@Router		/cart/{item_id}/{quantity} [put]
func CustomerPutQuantity(ctx *gin.Context) {
	item_id := ctx.Param("item_id")
	if item_id == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Invalid Item ID not set"))
		return
	}

	_quantity := ctx.Param("quantity")
	if _quantity == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Invalid quantity"))
		return
	}

	quantity, err := strconv.Atoi(_quantity)
	if err != nil {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Quantity must be an integer"))
		return
	}

	item := new(db.Item)
	if err := item.LoadFromDB(item_id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	quantity64 := int64(quantity)
	item.Quantity = &quantity64
	if err := validateItemQuantity(item, &item.Product); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	item.Amount = getItemAmount(&item.Product, *item.Quantity)

	if err := item.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Item updated successfully",
		"data": item,
	})
}

// CustomerDeleteFromCart Deletes an item with given id from the database.
//
//	@Summary	Allows the current user delete an item from the cart.
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		item_id			path		string					true	"The ID of the item to delete"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart/{item_id} [delete]
func CustomerDeleteFromCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	id := ctx.Param("item_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Item ID missing"))
		return
	}

	if err := db.DeleteItem(*user.ID, id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Item deleted successfully",
	})
}

// CustomerClearCart Clear the user's cart
//
//	@Summary	Allows the current user all items in their cart.
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart [delete]
func CustomerClearCart(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	if err := db.ClearCart(*user.ID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Cart cleared successfully",
	})
}
