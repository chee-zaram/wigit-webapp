package handlers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/chee-zaram/wigit-webapp/backend/internal/db"
	"github.com/gin-gonic/gin"
)

// ResetPasswordRequest is used to obtain the post and put data during password reset.
type ResetPasswordRequest struct {
	Email             string `json:"email" binding:"required,email,max=45"`
	NewPassword       string `json:"new_password" binding:"required,min=8,max=45"`
	RepeatNewPassword string `json:"repeat_new_password" binding:"required,min=8,max=45"`
	ResetToken        string `json:"reset_token" binding:"required"`
}

// cleanUp removes all leading and trailing whitespace from a string fields.
func (r *ResetPasswordRequest) cleanUp() {
	if r == nil {
		return
	}

	r.NewPassword = strings.TrimSpace(r.NewPassword)
	r.RepeatNewPassword = strings.TrimSpace(r.RepeatNewPassword)
	r.ResetToken = strings.TrimSpace(r.ResetToken)
}

// PostEmail binds to the post request body made to the `reset_password` endpoint.
type PostEmail struct {
	Email string `json:"email" binding:"required,email,max=45"`
}

// PostResetPassword Sends a request for a password update
//
//	@Summary	Allows to send a request for password update. A token is returned.
//	@Tags		reset_password
//	@Accept		json
//	@Produce	json
//	@Param		reset	body		PostEmail				true	"Email body"
//	@Success	201		{object}	map[string]interface{}	"reset_token"
//	@Failure	400		{object}	map[string]interface{}	"error"
//	@Failure	500		{object}	map[string]interface{}	"error"
//	@Router		/reset_password [post]
func PostResetPassword(ctx *gin.Context) {
	resetUser := new(PostEmail)

	if err := ctx.ShouldBind(resetUser); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	user := new(db.User)
	if code, err := user.LoadByEmail(resetUser.Email); err != nil {
		AbortCtx(ctx, code, err)
		return
	}

	randomBytes, err := generateRandomBytes()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}
	// token := base64.URLEncoding.EncodeToString(randomBytes)[:len(randomBytes)]
	token := base64.URLEncoding.EncodeToString(randomBytes)

	if err := user.UpdateResetToken(token); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"reset_token": user.ResetToken,
	})
}

// PutResetPassword Send a new password
//
//	@Summary	Allows to send new password details
//	@Tags		reset_password
//	@Accept		json
//	@Produce	json
//	@Param		reset	body		ResetPasswordRequest	true	"Data needed to reset the user's password"
//	@Success	200		{object}	map[string]interface{}	"msg"
//	@Failure	400		{object}	map[string]interface{}	"error"
//	@Failure	500		{object}	map[string]interface{}	"error"
//	@Router		/reset_password [put]
func PutResetPassword(ctx *gin.Context) {
	user, code, err := validateResetPasswordData(ctx)
	if err != nil {
		AbortCtx(ctx, code, err)
		return
	}

	passwordHash, salt, err := hashPassword(user)
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	user.HashedPassword = passwordHash
	user.Salt = salt
	user.ResetToken = ""

	if err := user.SaveToDB(); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Password has been reset successfully",
	})
}

// validateResetPasswordData validates the fields provided for the reset of a user's password.
// It returns a user from the database, an exit code, and an error if any.
func validateResetPasswordData(ctx *gin.Context) (*db.User, int, error) {
	resetPasswordRequest := new(ResetPasswordRequest)
	if err := ctx.ShouldBindJSON(resetPasswordRequest); err != nil {
		return nil, http.StatusBadRequest, errors.New("Failed to bind to ResetPasswordRequest")
	}

	resetPasswordRequest.cleanUp()
	user := new(db.User)
	if code, err := user.LoadByEmail(resetPasswordRequest.Email); err != nil {
		return nil, code, err
	}

	if user.ResetToken == "" {
		return nil, http.StatusUnauthorized, errors.New("Request to reset password first")
	}

	if resetPasswordRequest.ResetToken != user.ResetToken {
		return nil, http.StatusUnauthorized, errors.New("Provided token is different from reset token")
	}

	if resetPasswordRequest.RepeatNewPassword != resetPasswordRequest.NewPassword {
		return nil, http.StatusBadRequest, ErrPassMismatch
	}

	user.Password = &resetPasswordRequest.NewPassword

	return user, http.StatusOK, nil
}
