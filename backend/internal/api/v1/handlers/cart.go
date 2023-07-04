package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-ng/webapp/backend/internal/db"
)

// ItemRequest binds to the json body on post request to cart.
type ItemRequest struct {
	// ProductID is the id of the product the item represents.
	ProductID *string `json:"product_id" binding:"required"`
	// Quantity is the number of the product for this order.
	Quantity *int64 `json:"quantity" binding:"required"`
}

// GetCustomerCart Get all items in a user's cart
//
//	@Summary	Allows the customer retrieve all the items in their cart
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart [get]
func GetCustomerCart(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	cartItems, err := db.GetItemsInCart(*loggedInUser.ID)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": cartItems,
	})
}

// PostItemToCustomerCart Add a new item to cart
//
//	@Summary	Allows the current user to add a new item to the cart
//	@Tags		cart
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		item			body		ItemRequest				true	"A new item to add to cart"
//	@Success	201				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart [post]
func PostItemToCustomerCart(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	itemRequest := new(ItemRequest)
	if err := ctx.ShouldBindJSON(itemRequest); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	item := createCartItem(itemRequest)
	item.UserID = loggedInUser.ID
	product := new(db.Product)
	if err := product.LoadFromDB(*item.ProductID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	// Validate item quantity against product stock
	if err := validateItemQuantity(item, product); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	item.Amount = calculateCartItemAmount(product, *item.Quantity)

	if item.IsDuplicate() {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Duplicate item in cart"))
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

// createCartItem creates a new db.Item object from the ItemRequest.
func createCartItem(itemRequest *ItemRequest) *db.Item {
	item := new(db.Item)
	item.Quantity = itemRequest.Quantity
	item.ProductID = itemRequest.ProductID

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

// calculateCartItemAmount Computes amount for an item.
func calculateCartItemAmount(product *db.Product, quantity int64) *decimal.Decimal {
	amount := product.Price.Mul(decimal.NewFromInt(quantity))
	return &amount
}

// PutCartItemQuantity Changes the quantity of an item in the cart
//
//	@Summary	Allows a customer edit the quantity of an item in the cart
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Authorization token format 'Bearer <token>'"
//	@Param		item_id			path		string					true	"ID of the item"
//	@Param		quantity		path		string					true	"New item quantity"
//	@Success	200				{object}	map[string]interface{}	"msg,data"
//	@Success	400				{object}	map[string]interface{}	"error"
//	@Success	500				{object}	map[string]interface{}	"error"
//	@Router		/cart/{item_id}/{quantity} [put]
func PutCartItemQuantity(ctx *gin.Context) {
	itemID := ctx.Param("item_id")
	if itemID == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Invalid Item ID not set"))
		return
	}

	strItemQuantity := ctx.Param("quantity")
	if strItemQuantity == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Invalid quantity"))
		return
	}

	itemQuantity, err := strconv.Atoi(strItemQuantity)
	if err != nil {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Quantity must be an integer"))
		return
	}

	cartItem := new(db.Item)
	if err := cartItem.LoadFromDB(itemID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	itemQuantity64 := int64(itemQuantity)
	cartItem.Quantity = &itemQuantity64
	if err := validateItemQuantity(cartItem, &cartItem.Product); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	cartItem.Amount = calculateCartItemAmount(&cartItem.Product, *cartItem.Quantity)

	if err := cartItem.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Cart item quantity updated successfully",
		"data": cartItem,
	})
}

// DeleteItemFromCustomerCart Deletes an item with given id from the database.
//
//	@Summary	Delete an item from the customer's cart.
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		item_id			path		string					true	"The ID of the item to delete"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart/{item_id} [delete]
func DeleteItemFromCustomerCart(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	itemID := ctx.Param("item_id")
	if itemID == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Item ID missing"))
		return
	}

	if err := db.DeleteItem(*loggedInUser.ID, itemID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Item deleted successfully",
	})
}

// ClearCustomerCart Clear the user's cart
//
//	@Summary	Clears all the items in the logged in user's cart.
//	@Tags		cart
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/cart [delete]
func ClearCustomerCart(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	if err := db.ClearCart(*loggedInUser.ID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Cart cleared successfully",
	})
}
