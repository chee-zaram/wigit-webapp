package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	// ErrInternalServer indicates some server side error occured that can't be handled.
	ErrInternalServer = errors.New("Something went wrong!")
	// ErrDuplicateUser indicates that User with email already exists during signup.
	ErrDuplicateUser = errors.New("User with email already exists")
	// ErrInvalidPass indicates the user did not provide a Password on sign up or a valid password on signin.
	ErrInvalidPass = errors.New("Valid Password required")
	// ErrPassMismatch indicates the user entered did not repeat the password correctly.
	ErrPassMismatch = errors.New("Password Mismatch")
	// ErrInvalidProductID indicates that the id provided is invalid.
	ErrInvalidProductID = errors.New("Invalid Product ID")
	// ErrInvalidCategory indicates that the category of products provided is invalid.
	ErrInvalidCategory = errors.New("Not a valid category")
	// ErrInvalidServiceID indicates the service specified by the given ip does not exist.
	ErrInvalidServiceID = errors.New("Not a valid service")
	// ErrInvalidSlotID indicates that not slot id was provided or the slot id provided is not valid.
	ErrInvalidSlotID = errors.New("Slot ID is not valid")
	// ErrEmailParamNotSet indicates an email param is needed and was not provided.
	ErrEmailParamNotSet = errors.New("email param not set")
	// ErrUserCtx indicates the user is not set in the current context.
	ErrUserCtx = errors.New("User not set in context")
	// ErrStatusCtx indicates the status is not set in the current context.
	ErrStatusCtx = errors.New("Status not set")
)

// AbortCtx ends the current context with status and error message.
func AbortCtx(ctx *gin.Context, responseCode int, err error) {
	ctx.AbortWithStatusJSON(responseCode, gin.H{
		"error": err.Error(),
	})
}
