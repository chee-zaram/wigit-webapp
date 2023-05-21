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

// ConfigureJWT sets the JWTSigner and JWTVerifier using the secretKey
func ConfigureJWT(secretKey []byte) {
	// This may seem unnecessary, but using inferred types below will mean
	// the signer and verifier initialized will not be our global ones but
	// function scoped ones.
	var err error

	// Create a new signer for making tokens and building claims
	JWTSigner, err = jwt.NewSignerHS(jwt.HS256, secretKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create JWT signer")
	}

	// Create a new verifier validating token signatures
	JWTVerifier, err = jwt.NewVerifierHS(jwt.HS256, secretKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create JWT verifier")
	}
}

// CreateJWT returns a new token string set up for the logged in user using userID.
func CreateJWT(userID string) (string, error) {
	// Store user ID in claims
	claims := &jwt.RegisteredClaims{
		ID:        userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	builder := jwt.NewBuilder(JWTSigner)
	token, err := builder.Build(claims)
	if err != nil {
		log.Error().Err(err).Msg("failed to build token with claims")
		return "", err
	}

	return token.String(), nil
}
