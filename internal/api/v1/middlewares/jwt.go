package middlewares

import (
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/rs/zerolog/log"
)

var (
	JWTSigner   jwt.Signer
	JWTVerifier jwt.Verifier
)

// CreateSigner creates a new signer for making tokens and building claims.
func CreateSigner(secretKey []byte) {
	var err error
	JWTSigner, err = jwt.NewSignerHS(jwt.HS256, secretKey)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create JWT signer")
	}
}

// CreateVerifier creates a new verifier validating token signatures.
func CreateVerifier(secretKey []byte) {
	var err error
	JWTVerifier, err = jwt.NewVerifierHS(jwt.HS256, secretKey)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create JWT verifier")
	}
}

// CreateJWT returns a new token string set up for the logged in user using userID.
func CreateJWT(userID string) string {
	return buildWithClaims(generateClaims(userID))
}

// buildWithClaims returns a JWT token string built using given claims.
func buildWithClaims(claims *jwt.RegisteredClaims) string {
	builder := jwt.NewBuilder(JWTSigner)
	token, _ := builder.Build(claims)

	return token.String()
}

// generateClaims return registered claims with ID and expiration date set.
func generateClaims(userID string) *jwt.RegisteredClaims {
	// Store user ID in claims
	return &jwt.RegisteredClaims{
		ID:        userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 24)),
	}
}
