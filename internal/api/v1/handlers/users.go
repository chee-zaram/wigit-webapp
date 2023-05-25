package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// CustomerDeleteUser deletes the current user from the database.
func CustomerDeleteUser(ctx *gin.Context) {
	_, id, err := validateUserParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

// validateUserParams validates data sent to the `users` endpoint.
// It is used during updating and deletion of a user or information.
func validateUserParams(ctx *gin.Context) (*models.User, string, error) {
	_user, exists := ctx.Get("user")
	if !exists {
		return nil, "", errors.New("User not set in context")
	}
	user := _user.(*models.User)
	id := ctx.Param("user_id")
	if id == "" {
		return nil, "", errors.New("User ID must be set")
	}

	if id != *user.ID {
		return nil, "", errors.New("Cannot perform operation on another user's account")
	}

	return user, id, nil
}
