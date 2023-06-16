package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/wigit-gh/webapp/backend/internal/db"
	"gorm.io/gorm"
)

// NewProduct binds to the new product in the body of the request.
type NewProduct struct {
	Name        *string          `json:"name" binding:"required,min=3,max=45"`
	Description *string          `json:"description" binding:"required,min=3,max=1024"`
	Category    *string          `json:"category" binding:"required,min=3,max=45"`
	Stock       *int64           `json:"stock" binding:"required"`
	Price       *decimal.Decimal `json:"price" binding:"required"`
	ImageURL    *string          `json:"image_url" binding:"required,min=3,max=255"`
}

// validateData validates the fields provided in the json body during when adding new product.
func (product *NewProduct) validateData() error {
	if product.Price == nil || product.Price.Sign() < 0 {
		return errors.New("Invalid product price")
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
		AbortCtx(ctx, http.StatusBadRequest, errors.New("No product found"))
		return
	} else if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// GetProductsByCategory	Get products by category
//
//	@Summary	Retrieves a list of all products in given category
//	@Tags		products
//	@Produce	json
//	@Param		category	path		string					true	"Product Category"
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
		items, err := db.TrendingItems()
		if err != nil {
			AbortCtx(ctx, http.StatusInternalServerError, err)
			return
		}

		if products, err := db.TrendingProducts(items); err != nil {
			AbortCtx(ctx, http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": products})
		}
		return
	}

	products, err := db.ProductCategory(category)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

// AdminPostProducts	Post product
//
//	@Summary	Allows the admin add products to the database
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		product			body		NewProduct				true	"Add product"
//	@Success	201				{object}	map[string]interface{}	"data, msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/products [post]
func AdminPostProducts(ctx *gin.Context) {
	_newProduct := new(NewProduct)
	if err := ctx.ShouldBindJSON(_newProduct); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	if err := _newProduct.validateData(); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	product := newProduct(_newProduct)
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
func newProduct(newProduct *NewProduct) *db.Product {
	product := new(db.Product)
	product.Name = newProduct.Name
	product.Category = newProduct.Category
	product.Description = newProduct.Description
	product.ImageURL = newProduct.ImageURL
	product.Price = newProduct.Price
	product.Stock = newProduct.Stock

	return product
}

// AdminDeleteProducts	Deletes a product
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
func AdminDeleteProducts(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidProductID)
		return
	}

	if err := db.DeleteProduct(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Product delete successfully",
	})
}

// AdminPutProducts		Update product
//
//	@Summary	Allows the admin update the product with given product_id
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		product_id		path		string					true	"The id of the product to update"
//	@Param		product			body		NewProduct				true	"Update Product"
//	@Success	200				{object}	map[string]interface{}	"data, msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/products/{product_id} [put]
func AdminPutProducts(ctx *gin.Context) {
	_newProduct := new(NewProduct)
	id := ctx.Param("product_id")
	if id == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrInvalidProductID)
		return
	}

	if err := ctx.ShouldBindJSON(_newProduct); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	if err := _newProduct.validateData(); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	product := new(db.Product)
	if err := product.LoadFromDB(id); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	createdAt := product.CreatedAt
	product = newProduct(_newProduct)
	product.ID = &id
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
