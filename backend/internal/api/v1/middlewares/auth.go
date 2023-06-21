package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/backend/internal/db"
)

// JWTAuthentication validates a user's signin JWT set in the `Authorization` header.
func JWTAuthentication(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.Header(`WWW-Authenticate`, `Bearer realm="Restricted"`)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Missing Authorization header. Please provide a valid JWT",
		})
		return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", error="invalid_token", error_description="Invalid Authorization header format"`,
		))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Authorization header format. Please provide a valid JWT",
		})
		return
	}

	if bearerToken[0] != "Bearer" {
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", error="invalid_token", error_description="Authorization value does not contain Bearer"`,
		))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization value does not contain Bearer. Please provide a valid JWT",
		})
		return
	}

	userID, err := verifyJWTSignature(bearerToken[1])
	if err != nil {
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", error="invalid_token", error_description="%s"`, err.Error(),
		))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user := new(db.User)
	if err := user.LoadByID(userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}

// verifyJWTSignature checks the validity of the jwt provided.
// It returns the user ID stored in the claims, and any error if any occurs.
func verifyJWTSignature(_token string) (string, error) {
	token, err := parseToken(_token)
	if err != nil {
		return "", err
	}

	jwtClaims, err := retrieveTokenClaims(token)
	if err != nil {
		return "", err
	}

	if !jwtClaims.IsValidAt(time.Now().UTC()) {
		return "", errors.New("JWT has expired. Please sign in again.")
	}

	return jwtClaims.ID, nil
}

// parseToken takes a token as a string and verify the signature.
// It returns the parsed token as a pointer to a jwt.Token object.
func parseToken(_token string) (*jwt.Token, error) {
	token, err := jwt.Parse([]byte(_token), JWTVerifier)
	if err != nil {
		return nil, errors.New("failed to parse JWT. Please provide valid JWT.")
	}

	return token, nil
}

// retrieveTokenClaims return the claims stored in a token and any error.
func retrieveTokenClaims(token *jwt.Token) (*jwt.RegisteredClaims, error) {
	claims := new(jwt.RegisteredClaims)
	if err := json.Unmarshal(token.Claims(), claims); err != nil {
		return nil, errors.New("Failed to retrieve JWT claims. Please provide a valid JWT.")
	}

	return claims, nil
}

// AdminAuthorization validates if the user has admin privileges or not.
func AdminAuthorization(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in ctx"})
		return
	}

	if *loggedInUser.Role != "admin" && *loggedInUser.Role != "super_admin" {
		err := "You are not allowed to view this resource"
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", scope="admin super_admin", error="insufficient_scope", error_description="%s"`, err,
		))
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err})
		return
	}

	ctx.Next()
}

// SuperAdminAuthorization validates if the user is the super admin.
func SuperAdminAuthorization(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	loggedInUser, ok := userCtx.(*db.User)
	if !exists || !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not set in ctx"})
		return
	}

	if *loggedInUser.Role != "super_admin" {
		err := "You are not allowed to view this resource"
		ctx.Header(`WWW-Authenticate`, fmt.Sprintf(
			`Bearer realm="Restricted", scope="super_admin", error="insufficient_scope", error_description="%s"`, err,
		))
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err})
		return
	}

	ctx.Next()
}
