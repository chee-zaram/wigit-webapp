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
