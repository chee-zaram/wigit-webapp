package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	ginHost   = "WIGIT_GIN_HOST"
	ginPort   = "WIGIT_GIN_PORT"
	dbHost    = "WIGIT_DB_HOST"
	dbPort    = "WIGIT_DB_PORT"
	dbName    = "WIGIT_DB_NAME"
	dbUser    = "WIGIT_DB_USER"
	dbPass    = "WIGIT_DB_PASS"
	jwtSecret = "WIGIT_JWT_SECRET"
)

type Config struct {
	// GinHost is the host the gin app should connect to
	GinHost string
	// GinPort is the port the gin app should bind to, and listen and serve on
	GinPort string
	// DBHost is the network address or domain name of the server where the database is located
	DBHost string
	// DBPort is the communication endpoint on the server where the database is locoated
	DBPort string
	// DBName is the name of the database
	DBName string
	// DBUser is the user to connect to the database as
	DBUser string
	// DBPass is the password for the `DBUser`
	DBPass string
	// JWTSecret is the secret used to encrypt the jwt tokens
	JWTSecret string
	// Env is the environment where the server is currently running. Either `dev` or `prod`
	Env string
}

// NewConfig returns a configuration struct based on the current values of the
// environment variables, and the value of `env`
func NewConfig(env string) Config {
	return Config{
		GinHost:   getVariableValue(ginHost),
		GinPort:   getVariableValue(ginPort),
		DBHost:    getVariableValue(dbHost),
		DBPort:    getVariableValue(dbPort),
		DBName:    getVariableValue(dbName),
		DBUser:    getVariableValue(dbUser),
		DBPass:    getVariableValue(dbPass),
		JWTSecret: getVariableValue(jwtSecret),
		Env:       env,
	}
}

// getVariableValue returns the value of the environment variable `envVar`, or logs and panics
// if it is not set or invalid
func getVariableValue(envVar string) string {
	envVal, ok := os.LookupEnv(envVar)
	if !ok || envVal == "" {
		log.Panic().Str("var", envVar).Msg("environment variable not set or value is invalid")
	}

	if strings.Contains(envVar, "PORT") {
		if _, err := strconv.Atoi(envVal); err != nil {
			log.Panic().Str("var", envVar).Msg("value must be an integer")
		}
	}

	return envVal
}
