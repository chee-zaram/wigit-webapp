package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/backend/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/backend/internal/db"
	"golang.org/x/crypto/bcrypt"
)

// UserCredentials binds to the json body during signin.
type UserCredentials struct {
	Email    *string `json:"email" binding:"required,email,min=3,max=45"`
	Password *string `json:"password" binding:"required,min=8,max=45"`
}

// SignIn		Sign a user in
//
//	@Summary	Authenticate a user and generate a JWT.
//	@Tags		signin
//	@Accept		json
//	@Produce	json
//	@Param		user	body		UserCredentials			true	"Sign user in"
//	@Success	200		{object}	map[string]interface{}	"jwt, msg, user"
//	@Failure	400		{object}	map[string]interface{}	"error"
//	@Failure	500		{object}	map[string]interface{}	"error"
//	@Router		/signin [post]
func SignIn(ctx *gin.Context) {
	userCredentials := new(UserCredentials)
	if err := ctx.ShouldBind(userCredentials); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	user, code, err := verifyUserCredentials(userCredentials)
	if err != nil {
		AbortCtx(ctx, code, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Sign in successful",
		"user": user,
		"jwt":  middlewares.CreateJWT(*user.ID),
	})
}

// verifyUserCredentials loads user from database and validates the user credentials.
//
// It returns the user from the database and the code and error if any.
func verifyUserCredentials(user *UserCredentials) (*db.User, int, error) {
	// Get user with Email from the database.
	dbUser := new(db.User)
	if code, err := dbUser.LoadByEmail(*user.Email); err != nil {
		return nil, code, err
	}

	if err := validateUserPassword(user, dbUser); err != nil {
		return nil, http.StatusBadRequest, err
	}

	return dbUser, http.StatusOK, nil
}

// validateUserPassword verifies the user's password.
func validateUserPassword(user *UserCredentials, dbUser *db.User) error {
	salted := append([]byte(*user.Password), dbUser.Salt...)
	if err := bcrypt.CompareHashAndPassword(dbUser.HashedPassword, salted); err != nil {
		return ErrInvalidPass
	}
	return nil
}
