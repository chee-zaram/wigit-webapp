package handlers

import (
	"errors"
	"net/http"

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

	// If no products were found
	if products == nil {
		products = []models.Product{}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": products})
}

// GetProduct retrieves a product based on its id.
func GetProduct(ctx *gin.Context) {
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

	// If not record was found
	if products == nil {
		products = []models.Product{}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": products})
}

// getTrending finds all trending products from the database.
func getTrending() ([]models.Product, error) {
	panic("Not yet implemented")
}

// AdminPostProducts adds products to the database.
func AdminPostProducts(ctx *gin.Context) {
	product := new(models.Product)
	if err := ctx.ShouldBind(product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validatePostProductsData(product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Create(product).Error
	}); err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Product already exists"})
		return
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":     "Product created successfully",
		"product": product,
	})
}

// validatePostProductsData validates the fields provided in the json payload during
// post request to products endpoint.
func validatePostProductsData(product *models.Product) error {
	if product.Name == nil || *product.Name == "" {
		return errors.New("Invalid product name")
	}

	if product.Description == nil || *product.Description == "" {
		return errors.New("Invalid product description")
	}

	if product.Category == nil || *product.Category == "" {
		return errors.New("Invalid product category")
	}

	if product.Stock == nil || *product.Stock < 0 {
		return errors.New("Invalid product stock")
	}

	if product.Price == nil || product.Price.Sign() < 0 {
		return errors.New("Invalid product price")
	}

	if product.ImageURL == nil || *product.ImageURL == "" {
		return errors.New("Invalid image URL")
	}

	return nil
}
