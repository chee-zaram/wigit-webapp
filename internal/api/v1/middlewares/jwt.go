package middlewares

import (
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
