package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// CustomerDeleteUser deletes the current user from the database.
func CustomerDeleteUser(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
		return
	}
	user := _user.(*models.User)
	id := ctx.Param("user_id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID must be set"})
		return
	}

	if id != *user.ID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot delete another user's account"})
		return
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM users WHERE id = ?`, id).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User deleted successfully",
	})
}
