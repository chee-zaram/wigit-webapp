package middlewares

import (
	"testing"

	"github.com/cristalhq/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// TestCreateSigner tests the createSigner function with valid input.
func TestCreateSigner(t *testing.T) {
	assert := assert.New(t)
	JWTSigner = nil
	assert.Nil(JWTSigner)

	CreateSigner([]byte("jwtsecret"))
	signerType := new(jwt.HSAlg)

	assert.IsType(signerType, JWTSigner)
}

// TestCreateSigner_Invalid tests the CreateSigner function with invalid input.
func TestCreateSigner_Invalid(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() { CreateSigner([]byte("")) })
}

// TestCreateVerifier tests the CreateVerifier function with valid input.
func TestCreateVerifier(t *testing.T) {
	assert := assert.New(t)
	JWTVerifier = nil
	assert.Nil(JWTVerifier)

	CreateVerifier([]byte("jwtsecret"))
	verifierType := new(jwt.HSAlg)

	assert.IsType(verifierType, JWTVerifier)
}

// TestCreateVerifier_Invalid tests the CreateVerifier function with invalid input.
func TestCreateVerifier_Invalid(t *testing.T) {
	assert.Panics(t, func() { CreateVerifier([]byte("")) })
}

// TestGenerateClaims tests the generateClaims function.
func TestGenerateClaims(t *testing.T) {
	assert := assert.New(t)
	claims := generateJWTClaims("1234")
	assert.NotNil(claims)
	assert.Equal(claims.ID, "1234")
}

// TestBuildWithClaims tests the buildWithClaims function.
func TestBuildWithClaims(t *testing.T) {
	assert := assert.New(t)
	CreateSigner([]byte("jwtsecret"))
	token := buildJWTString(generateJWTClaims("1234"))
	assert.IsType("", token)
	assert.True(len(token) > 0)
}

// TestCreateJWT tests creation of jwt with valid input.
func TestCreateJWT(t *testing.T) {
	assert := assert.New(t)
	CreateSigner([]byte("jwtsecret"))
	CreateVerifier([]byte("jwtsecret"))
	token := CreateJWT("1234")
	assert.IsType("", token)
	assert.True(len(token) > 0)
}
