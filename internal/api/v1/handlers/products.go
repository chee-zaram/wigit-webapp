package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// GetProducts retrieves a list of all products.
func GetProducts(ctx *gin.Context) {
	var products []models.Product

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Find(&products).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

// GetProduct retrieves a product based on its id.
func GetProductByID(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidProductID.Error()})
		return
	}

	product := new(models.Product)

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(product, "id = ?", id).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No product found"})
		return
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductByCategory retrieves a list of all products in a given category.
func GetProductByCategory(ctx *gin.Context) {
	var products []models.Product

	category := ctx.Param("category")
	if category == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidCategory.Error()})
		return
	}

	if category == "trending" {
		if products, err := getTrending(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": products})
		}
		return
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Where("category = ?", category).Find(&products).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

// getTrending finds all trending products from the database.
func getTrending() ([]models.Product, error) {
	panic("Not yet implemented")
}

// AdminPostProducts adds products to the database.
func AdminPostProducts(ctx *gin.Context) {
	_product := new(models.Product)
	if err := ctx.ShouldBindJSON(_product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateProductsData(_product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Create(_product).Error
	}); err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Product already exists"})
		return
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	product, err := getProductFromDB(*_product.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Product created successfully",
		"data": product,
	})
}

// validateProductsData validates the fields provided in the json payload during
// post request to products endpoint.
func validateProductsData(product *models.Product) error {
	if product.Price == nil || product.Price.Sign() < 0 {
		return errors.New("Invalid product price")
	}

	return nil
}

// getProductFromDB retrieves a product from the database based the id.
func getProductFromDB(id string) (*models.Product, error) {
	product := new(models.Product)

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(product, "id = ?", id).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Product already exists")
	} else if err != nil {
		return nil, ErrInternalServer
	}

	return product, nil
}

// AdminDeleteProducts deletes a product from the database.
func AdminDeleteProducts(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidProductID.Error()})
		return
	}

	if err := deleteProductFromDB(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Product delete successfully",
	})
}

// deleteProductFromDB deletes a product from the database.
func deleteProductFromDB(id string) error {
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM products WHERE id = ?`, id).Error
	}); err != nil {
		return err
	}

	return nil
}

// AdminPutProducts updates the columns for the product with given id.
func AdminPutProducts(ctx *gin.Context) {
	_product := new(models.Product)
	id := ctx.Param("product_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidProductID.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(_product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := getProductFromDB(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := validateProductsData(_product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = updateProductInDB(product, _product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Product updated successfully",
		"data": product,
	})
}

// updateProductInDB updates a given product `dbProduct` with values from `newProduct`.
func updateProductInDB(dbProduct, newProduct *models.Product) error {
	dbProduct.Name = newProduct.Name
	dbProduct.Description = newProduct.Description
	dbProduct.Category = newProduct.Category
	dbProduct.Stock = newProduct.Stock
	dbProduct.Price = newProduct.Price
	dbProduct.ImageURL = newProduct.ImageURL

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Save(dbProduct).Error
	}); err != nil {
		return err
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(dbProduct, "id = ?", *dbProduct.ID).Error
	}); err != nil {
		return err
	}

	return nil
}
