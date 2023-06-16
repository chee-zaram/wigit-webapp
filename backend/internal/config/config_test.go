package config

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

// TestMain runs before any test in this package runs.
// It sets all environment variables to ensure all tests can be carried out,
// and resets the variables when tests are done.
func TestMain(m *testing.M) {
	oldGinHostVal := os.Getenv(ginHost)
	if err := os.Setenv(ginHost, "localhost"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(ginHost, oldGinHostVal)

	oldGinPortVal := os.Getenv(ginPort)
	if err := os.Setenv(ginPort, "5001"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(ginPort, oldGinPortVal)

	oldDBHostVal := os.Getenv(dbHost)
	if err := os.Setenv(dbHost, "localhost"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(dbHost, oldDBHostVal)

	oldDBPortVal := os.Getenv(dbPort)
	if err := os.Setenv(dbPort, "3306"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(dbPort, oldDBPortVal)

	oldDBNameVal := os.Getenv(dbName)
	if err := os.Setenv(dbName, "wigit_webapp_test_db"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(dbName, oldDBNameVal)

	oldDBUserVal := os.Getenv(dbUser)
	if err := os.Setenv(dbUser, "wigit_webapp_user"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(dbUser, oldDBUserVal)

	oldDBPassVal := os.Getenv(dbPass)
	if err := os.Setenv(dbPass, "wigit_webapp_pwd"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(dbPass, oldDBPassVal)

	oldJWTValue := os.Getenv(jwtSecret)
	if err := os.Setenv(jwtSecret, "wigit_jwt_secret"); err != nil {
		log.Panic().Err(err).Msg("failed to set env in config_test")
	}
	defer os.Setenv(jwtSecret, oldJWTValue)

	exitCode := m.Run()
	os.Exit(exitCode)
}

// TestNewDevConfig tests that all variables are set and returned.
func TestNewDevConfig(t *testing.T) {
	assert := assert.New(t)
	conf := NewConfig("dev")
	assert.NotEqual("", conf.GinHost)
	assert.NotEqual("", conf.GinPort)
	assert.NotEqual("", conf.DBHost)
	assert.NotEqual("", conf.DBPort)
	assert.NotEqual("", conf.DBName)
	assert.NotEqual("", conf.DBPass)
	assert.Equal("dev", conf.Env)
}

// TestNewProdConfig tests the configuration settings for production.
func TestNewProdConfig(t *testing.T) {
	assert := assert.New(t)
	conf := NewConfig("prod")
	assert.NotEqual("", conf.GinHost)
	assert.NotEqual("", conf.GinPort)
	assert.NotEqual("", conf.DBHost)
	assert.NotEqual("", conf.DBPort)
	assert.NotEqual("", conf.DBName)
	assert.NotEqual("", conf.DBPass)
	assert.Equal("prod", conf.Env)
}

// TestNewConfigHostNotSet this checks that the program panics when a host is not set
// for the gin app.
func TestNewConfigHostNotSet(t *testing.T) {
	assert := assert.New(t)
	host, ok := os.LookupEnv(ginHost)
	err := os.Setenv(ginHost, "")
	defer os.Setenv(ginHost, host)

	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}

// TestNewConfigPortNotSet checks that the program panics when no port is
// set for the gin app.
func TestNewConfigPortNotSet(t *testing.T) {
	assert := assert.New(t)
	port, ok := os.LookupEnv(ginPort)
	err := os.Setenv(ginPort, "")
	defer os.Setenv(ginPort, port)

	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}

// TestNewConfigPortNotValid checks that the program panics when an invalid port number
// is given for the gin app.
func TestNewConfigPortNotValid(t *testing.T) {
	assert := assert.New(t)
	port, ok := os.LookupEnv(ginPort)
	err := os.Setenv(ginPort, "this is not a port number of an integer")
	defer os.Setenv(ginPort, port)
	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}

// TestNewConfigDBHostNotSet checks that the program panics when no Host is set for
// the database.
func TestNewConfigDBHostNotSet(t *testing.T) {
	assert := assert.New(t)
	_dbHost, ok := os.LookupEnv(dbHost)
	err := os.Setenv(dbHost, "")
	defer os.Setenv(dbHost, _dbHost)

	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}

// TestNewConfigDBPortNotSet checks that the program panics when no port is set for
// the database.
func TestNewConfigDBPortNotSet(t *testing.T) {
	assert := assert.New(t)
	_dbPort, ok := os.LookupEnv(dbPort)
	err := os.Setenv(dbPort, "")
	defer os.Setenv(dbPort, _dbPort)

	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}

// TestNewConfigDBPortNotValid checks that the program panics when an invalid port
// number is specified for the database.
func TestNewConfigDBPortNotValid(t *testing.T) {
	assert := assert.New(t)
	_dbPort, ok := os.LookupEnv(dbPort)
	err := os.Setenv(dbPort, "this is not a valid port number or an integer")
	defer os.Setenv(dbPort, _dbPort)

	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}

// TestNewConfigDBNameNotSet checks that the program panics when no database name is given.
func TestNewConfigDBNameNotSet(t *testing.T) {
	assert := assert.New(t)
	_dbName, ok := os.LookupEnv(dbName)
	err := os.Setenv(dbName, "")
	defer os.Setenv(dbName, _dbName)

	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}

// TestNewConfigDBPasswordNotSet checks that the program panics when no password
// is set for the database user.
func TestNewConfigDBPasswordNotSet(t *testing.T) {
	assert := assert.New(t)
	dbPassword, ok := os.LookupEnv(dbPass)
	err := os.Setenv(dbPass, "")
	defer os.Setenv(dbPass, dbPassword)

	assert.True(ok)
	assert.Nil(err)
	assert.Panics(func() { NewConfig("dev") })
}
