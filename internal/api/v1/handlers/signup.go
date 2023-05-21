package handlers

import (
	"crypto/rand"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/wigit-gh/webapp/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/internal/db/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SignUp is the handler function for signing up users to the app.
func SignUp(ctx *gin.Context) {
	user := new(models.User)
	if err := ctx.ShouldBind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Sign Up failed."})
		return
	}

	if err := addUser(user); err != nil {
		log.Error().Err(err).Msg("")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := middlewares.CreateJWT(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "Sign up successful",
		"jwt": token,
	})
}

// addUser adds a new user to the database. Returns an error if any exists.
func addUser(user *models.User) error {
	if err := validateSignUpUser(user); err != nil {
		return err
	}

	if err := hashPassword(user); err != nil {
		return err
	}

	// Add the user to the database
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	}); err != nil {
		return ErrFailedToAddUserToDB
	}

	return nil
}

// validateSignUpUser validates all fields in the post form
func validateSignUpUser(user *models.User) error {
	var dbUser *models.User

	// Verify user does not already exist
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.First(dbUser, "email = ?", user.Email).Error
	}); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternalServer
	}
	if dbUser != nil && dbUser.Email != nil {
		return ErrDuplicateUser
	}

	if user.FirstName == nil {
		return ErrInvalidFirstName
	}

	if user.LastName == nil {
		return ErrInvalidLastName
	}

	if user.Address == nil {
		return ErrInvalidAddress
	}

	if user.Phone == nil {
		return ErrNoPhone
	}

	if len(*user.Phone) < 9 || len(*user.Phone) > 11 {
		return ErrInvalidPhone
	}

	if user.Password == nil || user.RepeatPassword == nil {
		return ErrInvalidPass
	}

	if *user.Password != *user.RepeatPassword {
		return ErrPassMismatch
	}

	if len(*user.RepeatPassword) < 8 {
		return ErrPassTooShort
	}

	if len(*user.Password) > 45 {
		return ErrPassTooLong
	}

	return nil
}

// hashPassword creates a hash of the password plus a random salt.
func hashPassword(user *models.User) error {
	salt, err := generateSalt()
	if err != nil {
		return err
	}

	passHash, err := createHash([]byte(*user.Password), salt)
	if err != nil {
		return err
	}

	user.HashedPassword = passHash
	user.Salt = salt

	return nil
}

// generateSalt creates a new random salt for the new user.
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)

	// generate random bytes
	if _, err := rand.Read(salt); err != nil {
		log.Error().Err(err).Msg("failed to create salt")
		return nil, err
	}
	return salt, nil
}

// createHash create the encrypted hashed password using bcrypt algorithm.
func createHash(password, salt []byte) ([]byte, error) {
	// Add the random salt to the password
	passPlusSalt := append(password, salt...)
	// Generate the hash
	passHash, err := bcrypt.GenerateFromPassword(passPlusSalt, bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate hashed password")
		return nil, err
	}

	return passHash, nil
}
