package handlers

import (
	"crypto/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/wigit-gh/webapp/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/internal/db"
	"golang.org/x/crypto/bcrypt"
)

// SignUpUser binds to a user during signup.
type SignUpUser struct {
	FirstName      *string `json:"first_name" binding:"required,min=3,max=45"`
	LastName       *string `json:"last_name" binding:"required,min=3,max=45"`
	Email          *string `json:"email" binding:"required,email,min=5,max=45"`
	Password       *string `json:"password" binding:"required,min=8,max=45"`
	RepeatPassword *string `json:"repeat_password" binding:"required,min=8,max=45"`
	Address        *string `json:"address" binding:"required,min=3,max=255"`
	Phone          *string `json:"phone" binding:"required,min=8,max=11"`
}

// validateSignUpUser validates all fields in the post request.
func (user *SignUpUser) validate() error {
	dbUser := new(db.User)
	if code, err := dbUser.LoadByEmail(*user.Email); code == http.StatusBadRequest {
	} else if dbUser.Email != nil {
		return ErrDuplicateUser
	} else if err != nil {
		return err
	}

	if *user.Password != *user.RepeatPassword {
		return ErrPassMismatch
	}

	return nil
}

// SignUp		Sign up a user
//
//	@Summary	Add a new user account
//	@Tags		signup
//	@Accept		json
//	@Produce	json
//	@Param		user	body		SignUpUser				true	"Add user account"
//	@Success	201		{object}	map[string]interface{}	"jwt, msg"
//	@Failure	400		{object}	map[string]interface{}	"error"
//	@Failure	500		{object}	map[string]interface{}	"error"
//	@Router		/signup [post]
func SignUp(ctx *gin.Context) {
	signUpUser := new(SignUpUser)
	if err := ctx.ShouldBind(signUpUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := signUpUser.validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, code, err := newUser(signUpUser)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	if err := user.SaveToDB(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "Sign up successful",
		"jwt": middlewares.CreateJWT(*user.ID),
	})
}

// newUser fills up the initial user fields from the SignUpUser struct.
// It returns a new user object, a status code, and an error if any.
func newUser(signUpUser *SignUpUser) (*db.User, int, error) {
	user := new(db.User)
	user.Email = signUpUser.Email
	user.Password = signUpUser.Password
	user.RepeatPassword = signUpUser.RepeatPassword
	user.FirstName = signUpUser.FirstName
	user.LastName = signUpUser.LastName
	user.Address = signUpUser.Address
	user.Phone = signUpUser.Phone

	passHash, salt, err := hashPassword(user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	user.HashedPassword = passHash
	user.Salt = salt

	return user, http.StatusCreated, nil
}

// hashPassword creates a hash of the password plus a random salt.
func hashPassword(user *db.User) (passHash []byte, salt []byte, err error) {
	salt, err = generateRandomBytes()
	if err != nil {
		return nil, nil, err
	}

	passHash, err = createHash([]byte(*user.Password), salt)
	if err != nil {
		return nil, nil, err
	}

	return passHash, salt, nil
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
