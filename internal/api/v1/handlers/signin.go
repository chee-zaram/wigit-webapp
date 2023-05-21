package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/internal/db/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SignIn handles post requests to the /signin route.
func SignIn(ctx *gin.Context) {
	user := new(models.User)
	if err := ctx.ShouldBind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := authenticateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := middlewares.CreateJWT(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Sign in successful",
		"jwt": token,
	})
}

// authenticateUser verifies the user attempting to log in is a valid user.
func authenticateUser(user *models.User) (string, error) {
	var err error

	if user.Email == nil {
		return "", ErrEmailNotProvided
	}

	if user.Password == nil {
		return "", ErrInvalidPass
	}

	// Get user with Email from the database.
	dbUser, err := getUserFromDB(*user.Email)
	if err != nil {
		return "", err
	}

	// Verify the user password.
	if err := validateUser(user, dbUser); err != nil {
		return "", err
	}

	return *dbUser.ID, nil
}

// getUserFromDB gets the user with `email` from the database.
func getUserFromDB(email string) (*models.User, error) {
	dbUser := new(models.User)

	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(dbUser, "email = ?", email).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidUser
	} else if err != nil {
		return nil, ErrInternalServer
	}

	return dbUser, nil
}

// validateUser verifies the user's password.
func validateUser(user, dbUser *models.User) error {
	salted := append([]byte(*user.Password), dbUser.Salt...)
	if err := bcrypt.CompareHashAndPassword(dbUser.HashedPassword, salted); err != nil {
		return ErrInvalidPass
	}
	return nil
}
