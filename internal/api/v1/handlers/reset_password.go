package handlers

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db"
	"github.com/wigit-gh/webapp/internal/db/models"
	"gorm.io/gorm"
)

// ResetPassword is used to obtain the post and put data during password reset.
type ResetPassword struct {
	Email             string `json:"email" binding:"required,email,min=5,max=45"`
	NewPassword       string `json:"new_password"`
	RepeatNewPassword string `json:"repeat_new_password"`
	ResetToken        string `json:"reset_token"`
}

// PostResetPassword sends a reset password request with user email.
func PostResetPassword(ctx *gin.Context) {
	resetUser := new(ResetPassword)

	if err := ctx.ShouldBind(resetUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrEmailNotProvided.Error()})
		return
	}

	user, code, err := getUserFromDB(resetUser.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	_token, err := generateRandomBytes()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token := base64.URLEncoding.EncodeToString(_token)[:len(_token)]

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Model(user).Update("reset_token", token).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"reset_token": token})
}

// PutResetPassword updates the user's password.
func PutResetPassword(ctx *gin.Context) {
	user, code, err := validateResetPasswordPutData(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	if err := hashPassword(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	user.ResetToken = ""
	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Password has been reset successfully"})
}

// validateResetPasswordPutData validates the fields provided for the reset of a user's password.
// It returns a user from the database, an exit code, and an error if any.
func validateResetPasswordPutData(ctx *gin.Context) (*models.User, int, error) {
	resetUser := new(ResetPassword)
	if err := ctx.ShouldBindJSON(resetUser); err != nil {
		return nil, http.StatusBadRequest, errors.New("Failed to bind")
	}

	if resetUser.ResetToken == "" {
		return nil, http.StatusBadRequest, errors.New("Reset Token not provided")
	}

	user, code, err := getUserFromDB(resetUser.Email)
	if err != nil {
		return nil, code, err
	}

	if user.ResetToken == "" {
		return nil, http.StatusUnauthorized, errors.New("Request to reset password first")
	}

	if resetUser.ResetToken != user.ResetToken {
		return nil, http.StatusUnauthorized, errors.New("Provided token is different from reset token")
	}

	if resetUser.RepeatNewPassword != resetUser.NewPassword {
		return nil, http.StatusBadRequest, ErrPassMismatch
	}

	if len(resetUser.NewPassword) < 8 {
		return nil, http.StatusBadRequest, ErrPassTooShort
	}

	if len(resetUser.NewPassword) > 45 {
		return nil, http.StatusBadRequest, ErrPassTooLong
	}

	user.Password = &resetUser.NewPassword

	return user, http.StatusOK, nil
}
