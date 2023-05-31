package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/internal/db/models"
	"golang.org/x/crypto/bcrypt"
)

// SignIn handles post requests to the /signin route.
func SignIn(ctx *gin.Context) {
	_user := new(models.User)
	if err := ctx.ShouldBind(_user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, code, err := authenticateUser(_user)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Sign in successful",
		"user": user,
		"jwt":  middlewares.CreateJWT(*user.ID),
	})
}

// authenticateUser verifies the user attempting to log in is a valid user.
func authenticateUser(user *models.User) (*models.User, int, error) {
	var err error

	if user.Email == nil {
		return nil, http.StatusBadRequest, ErrEmailNotProvided
	}

	if user.Password == nil {
		return nil, http.StatusBadRequest, ErrInvalidPass
	}

	// Get user with Email from the database.
	dbUser, code, err := getUserFromDB(*user.Email)
	if err != nil {
		return nil, code, err
	}

	// Verify the user password.
	if err := validateUser(user, dbUser); err != nil {
		return nil, code, err
	}

	return dbUser, code, nil
}

// validateUser verifies the user's password.
func validateUser(user, dbUser *models.User) error {
	salted := append([]byte(*user.Password), dbUser.Salt...)
	if err := bcrypt.CompareHashAndPassword(dbUser.HashedPassword, salted); err != nil {
		return ErrInvalidPass
	}
	return nil
}
