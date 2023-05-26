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
	user, _, err := validateUserParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM users WHERE id = ?`, *user.ID).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User deleted successfully",
	})
}

// CustomerPutUser updates a user's information in the database.
func CustomerPutUser(ctx *gin.Context) {
	user, newUser, err := validateUserParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUserInfo(user, newUser)
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := getUserByID(*user.ID)
	if err != nil {
		updatedUser = user
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User updated successfully",
		"data": updatedUser,
	})
}

// validateUserParams validates data sent to the `users` endpoint.
// It is used during updating and deletion of a user or information.
func validateUserParams(ctx *gin.Context) (*models.User, *models.User, error) {
	_user, exists := ctx.Get("user")
	if !exists {
		return nil, nil, errors.New("User not set in context")
	}
	user := _user.(*models.User)
	id := ctx.Param("user_id")
	if id == "" {
		return nil, nil, errors.New("User ID must be set")
	}

	if id != *user.ID {
		return nil, nil, errors.New("Cannot perform operation on another user's account")
	}

	newUser := new(models.User)
	if err := ctx.ShouldBind(newUser); err != nil {
		return nil, nil, err
	}

	return user, newUser, nil
}

// updateUserInfo updates a user's information on put request.
func updateUserInfo(user, newUser *models.User) {
	user.Email = newUser.Email

	if newUser.FirstName != nil && *newUser.FirstName != "" {
		user.FirstName = newUser.FirstName
	}

	if newUser.LastName != nil && *newUser.LastName != "" {
		user.LastName = newUser.LastName
	}

	if newUser.Address != nil && *newUser.Address != "" {
		user.Address = newUser.Address
	}

	if newUser.Phone != nil && *newUser.Phone != "" {
		user.Phone = newUser.Phone
	}
}
