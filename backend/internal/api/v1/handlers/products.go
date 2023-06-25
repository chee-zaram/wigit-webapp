package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/backend/internal/db"
	"gorm.io/gorm"
)

// ProductRequest binds to the new product in the body of the request.
type ProductRequest struct {
	Name        *string          `json:"name" binding:"required,min=3,max=45"`
	Description *string          `json:"description" binding:"required,min=3,max=1024"`
	Category    *string          `json:"category" binding:"required,min=3,max=45"`
	Stock       *int64           `json:"stock" binding:"required"`
	Price       *decimal.Decimal `json:"price" binding:"required"`
	ImageURL    *string          `json:"image_url" binding:"required,min=3,max=255"`
}

// cleanUp removes all leading and trailing spaces from the data string fields.
func (p *ProductRequest) cleanUp() {
	if p == nil {
		return
	}

	*p.Name = strings.TrimSpace(*p.Name)
	*p.Description = strings.TrimSpace(*p.Description)
	*p.Category = strings.TrimSpace(*p.Category)
	*p.ImageURL = strings.TrimSpace(*p.ImageURL)
}

// validateData validates the fields provided in the json body during when adding new product.
func (product *ProductRequest) validateData() error {
	if product.Price == nil || product.Price.Sign() < 0 {
		return errors.New("Product price must be a positive decimal")
	}

	return nil
}

// GetProducts	Gets a list of all products
//
//	@Summary	Retrieves a list of all product objects
//	@Tags		products
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}	"data"
//	@Failure	500	{object}	map[string]interface{}	"error"
//	@Router		/products [get]
func GetProducts(ctx *gin.Context) {
	products, err := db.AllProducts()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

// GetProductByID	Gets a product by ID
//
//	@Summary	Retrieves the product with the given ID
//	@Tags		products
//	@Produce	json
//	@Param		product_id	path		string					true	"Product ID"
//	@Success	200			{object}	map[string]interface{}	"data"
//	@Failure	400			{object}	map[string]interface{}	"error"
//	@Failure	500			{object}	map[string]interface{}	"error"
//	@Router		/products/{product_id} [get]
func GetProductByID(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidProductID)
		return
	}

	product := new(db.Product)
	if err := product.LoadFromDB(id); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Product with ID not found in database"))
		return
	} else if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// GetProductsByName	Get products with the given substring in name.
//
//	@Summary	Retrieves a list of all products in given substring `name` in their name.
//	@Tags		products
//	@Produce	json
//	@Param		name	path		string					true	"The name to search for"
//	@Success	200		{object}	map[string]interface{}	"data"
//	@Failure	400		{object}	map[string]interface{}	"error"
//	@Failure	500		{object}	map[string]interface{}	"error"
//	@Router		/products/search/{name} [get]
func GetProductsByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Product name not set in cxt"))
		return
	}

	name = strings.ToLower(name)
	products, err := db.GetProductsByName(name)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, errors.New("Failed to get products by name"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

// GetProductsByCategory	Get products by category
//
//	@Summary	Retrieves a list of all products in given category
//	@Tags		products
//	@Produce	json
//	@Param		category	path		string					true	"The category of the product to retrieve"
//	@Success	200			{object}	map[string]interface{}	"data"
//	@Failure	400			{object}	map[string]interface{}	"error"
//	@Failure	500			{object}	map[string]interface{}	"error"
//	@Router		/products/categories/{category} [get]
func GetProductsByCategory(ctx *gin.Context) {
	var products []db.Product

	category := ctx.Param("category")
	if category == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidCategory)
		return
	}

	if category == "trending" {
		trendingItems, err := db.TrendingItems()
		if err != nil {
			AbortCtx(ctx, http.StatusInternalServerError, err)
			return
		}

		if products, err := db.TrendingProducts(trendingItems); err != nil {
			AbortCtx(ctx, http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": products})
		}
		return
	}

	products, err := db.GetProductsByCategory(category)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

// AdminPostProduct	Add product by admin
//
//	@Summary	Allows the admin add products to the database
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Authorization token format 'Bearer <token>'"
//	@Param		product			body		ProductRequest			true	"Add product"
//	@Success	201				{object}	map[string]interface{}	"data, msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/products [post]
func AdminPostProduct(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	admin, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	productRequest := new(ProductRequest)
	if err := ctx.ShouldBindJSON(productRequest); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	productRequest.cleanUp()
	if err := productRequest.validateData(); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	adminFullName := fmt.Sprintf("%s %s", *admin.FirstName, *admin.LastName)
	product := newProduct(productRequest, adminFullName, false)
	if err := product.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Product created successfully",
		"data": product,
	})
}

// newProduct gets a pointer to a new db.Product object.
func newProduct(newProduct *ProductRequest, adminName string, exists bool) *db.Product {
	product := new(db.Product)
	product.Name = newProduct.Name
	product.Category = newProduct.Category
	product.Description = newProduct.Description
	product.ImageURL = newProduct.ImageURL
	product.Price = newProduct.Price
	product.Stock = newProduct.Stock

	if !exists {
		product.AddedBy = adminName
		msg := fmt.Sprintf("product with name = [%s] added by [%s]", *product.Name, adminName)
		log.Info().Msg(msg)
	} else {
		product.UpdatedBy = adminName
		msg := fmt.Sprintf(
			"product with name = [%s] updated by [%s]. stock = %d and price = %s",
			*product.Name, adminName, *product.Stock, *product.Price,
		)
		log.Info().Msg(msg)
	}

	return product
}

// AdminDeleteProduct	Deletes a product by an admin
//
//	@Summary	Allows admins delete a product from the database
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		product_id		path		string					true	"Product id to delete"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/products/{product_id} [delete]
func AdminDeleteProduct(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	admin, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	id := ctx.Param("product_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidProductID)
		return
	}

	adminName := fmt.Sprintf("%s %s", *admin.FirstName, *admin.LastName)
	if err := db.DeleteProduct(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}
	log.Info().Msg(fmt.Sprintf("product [%s] deleted by [%s]", id, adminName))

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Product delete successfully",
	})
}

// AdminPutProduct		Update product
//
//	@Summary	Allows the admin update the product with given product_id
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		product_id		path		string					true	"The id of the product to update"
//	@Param		product			body		ProductRequest			true	"Request body containing the new product data"
//	@Success	200				{object}	map[string]interface{}	"data, msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/products/{product_id} [put]
func AdminPutProduct(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	admin, ok := userCtx.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	productRequest := new(ProductRequest)
	productID := ctx.Param("product_id")
	if productID == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidProductID)
		return
	}

	if err := ctx.ShouldBindJSON(productRequest); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	productRequest.cleanUp()
	if err := productRequest.validateData(); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	product := new(db.Product)
	if err := product.LoadFromDB(productID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	createdAt := product.CreatedAt
	adminFullName := fmt.Sprintf("%s %s", *admin.FirstName, *admin.LastName)
	product = newProduct(productRequest, adminFullName, true)
	product.ID = &productID
	product.CreatedAt = createdAt

	if err := product.Reload(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Product updated successfully",
		"data": product,
	})
}
