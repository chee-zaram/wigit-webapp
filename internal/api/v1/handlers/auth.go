package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/internal/db"
)

// JWTAuthentication validates a user's signin JWT token set in the `Authorization` header.
func JWTAuthentication(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.Header(`WWW-Authenticate`, `Bearer realm="Restricted"`)
		AbortCtx(ctx, http.StatusUnauthorized, errors.New("Authorization header missing"))
		return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", error="invalid_token", error_description="Invalid Authorization header format"`,
		))
		AbortCtx(ctx, http.StatusUnauthorized, errors.New("Invalid Authorization header format"))
		return
	}

	if bearerToken[0] != "Bearer" {
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", error="invalid_token", error_description="Authorization value does not contain Bearer"`,
		))
		AbortCtx(ctx, http.StatusUnauthorized, errors.New("Authorization value does not contain Bearer"))
		return
	}

	userID, err := validateJWTToken(bearerToken[1])
	if err != nil {
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", error="invalid_token", error_description="%s"`, err.Error(),
		))
		AbortCtx(ctx, http.StatusUnauthorized, err)
		return
	}

	user := new(db.User)
	if err := user.LoadByID(userID); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}

// validateJWTToken checks the validity of the jwt token provided.
// It returns the user ID stored in the claims, and any error if any occurs.
func validateJWTToken(_token string) (string, error) {
	token, err := parseToken(_token)
	if err != nil {
		return "", err
	}

	claims, err := retrieveTokenClaims(token)
	if err != nil {
		return "", err
	}

	if !claims.IsValidAt(time.Now().UTC()) {
		return "", errors.New("Token has expired")
	}

	return claims.ID, nil
}

// parseToken takes a token as a string and verify the signature.
// It returns the parsed token as a pointer to a jwt.Token object.
func parseToken(_token string) (*jwt.Token, error) {
	token, err := jwt.Parse([]byte(_token), middlewares.JWTVerifier)
	if err != nil {
		return nil, errors.New("failed to parse JWT token")
	}

	return token, nil
}

// retrieveTokenClaims return the claims stored in a token and any error.
func retrieveTokenClaims(token *jwt.Token) (*jwt.RegisteredClaims, error) {
	claims := new(jwt.RegisteredClaims)
	if err := json.Unmarshal(token.Claims(), claims); err != nil {
		return nil, errors.New("failed to Unmarshal claims")
	}

	return claims, nil
}

// Authorization validates if the user has admin privileges or not.
func AdminAuthorization(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	if *user.Role != "admin" && *user.Role != "super_admin" {
		err := "You are not allowed to view this resource"
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", scope="admin super_admin", error="insufficient_scope", error_description="%s"`, err,
		))
		AbortCtx(ctx, http.StatusForbidden, errors.New(err))
		return
	}

	ctx.Next()
}

// SuperAdminAuthorization validates if the user is the super admin.
func SuperAdminAuthorization(ctx *gin.Context) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		AbortCtx(ctx, http.StatusBadRequest, ErrUserCtx)
		return
	}

	if *user.Role != "super_admin" {
		err := "You are not allowed to view this resource"
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", scope="super_admin", error="insufficient_scope", error_description="%s"`, err,
		))
		AbortCtx(ctx, http.StatusForbidden, errors.New(err))
		return
	}

	ctx.Next()
}
