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

	if code, err := addUser(user); err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "Sign up successful",
		"jwt": middlewares.CreateJWT(*user.ID),
	})
}

// addUser adds a new user to the database. Returns an error if any exists.
func addUser(user *models.User) (int, error) {
	if err := validateSignUpUser(user); err != nil {
		return http.StatusBadRequest, err
	}

	if err := hashPassword(user); err != nil {
		return http.StatusInternalServerError, err
	}

	// Add the user to the database
	if err := DBConnector.Query(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	}); err != nil {
		return http.StatusInternalServerError, ErrFailedToAddUserToDB
	}

	return http.StatusCreated, nil
}

// validateSignUpUser validates all fields in the post form
func validateSignUpUser(user *models.User) error {
	// Verify user does not already exist
	if dbUser, _, err := getUserFromDB(*user.Email); errors.Is(err, ErrInvalidUser) {
	} else if dbUser != nil {
		return ErrDuplicateUser
	} else if err != nil {
		return err
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

	// We don't want a user's role to be set during signup but only through
	// the database.
	user.Role = nil

	return nil
}

// hashPassword creates a hash of the password plus a random salt.
func hashPassword(user *models.User) error {
	salt, err := generateRandomBytes()
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

// generateRandomBytes generates a random 16 byte slice.
func generateRandomBytes() ([]byte, error) {
	salt := make([]byte, 16)

	// generate random bytes
	if _, err := rand.Read(salt); err != nil {
		log.Error().Err(err).Msg("failed to create salt")
		return nil, ErrInternalServer
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
