package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/internal/db"
	"golang.org/x/crypto/bcrypt"
)

// SignInUser binds to the json body during signin.
type SignInUser struct {
	// Email is the user's email.
	Email *string `json:"email" binding:"required,email,min=3,max=45"`
	// Password is the user's password.
	Password *string `json:"password" binding:"required,min=8,max=45"`
}

// SignIn		Sign a user in
//
//	@Summary	Logs a user into their account
//	@Tags		signin
//	@Accept		json
//	@Produce	json
//	@Param		user	body		SignInUser				true	"Sign user in"
//	@Success	200		{object}	map[string]interface{}	"jwt, msg, user"
//	@Failure	400		{object}	map[string]interface{}	"error"
//	@Failure	500		{object}	map[string]interface{}	"error"
//	@Router		/signin [post]
func SignIn(ctx *gin.Context) {
	_user := new(SignInUser)
	if err := ctx.ShouldBind(_user); err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	user, code, err := authenticateUser(_user)
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

// authenticateUser verifies the user attempting to log in is a valid user.
func authenticateUser(user *SignInUser) (*db.User, int, error) {
	// Get user with Email from the database.
	dbUser := new(db.User)
	if code, err := dbUser.LoadByEmail(*user.Email); err != nil {
		return nil, code, err
	}

	// Verify the user password.
	if err := validateUser(user, dbUser); err != nil {
		return nil, http.StatusBadRequest, err
	}

	return dbUser, http.StatusOK, nil
}

// validateUser verifies the user's password.
func validateUser(user *SignInUser, dbUser *db.User) error {
	salted := append([]byte(*user.Password), dbUser.Salt...)
	if err := bcrypt.CompareHashAndPassword(dbUser.HashedPassword, salted); err != nil {
		return ErrInvalidPass
	}
	return nil
}
