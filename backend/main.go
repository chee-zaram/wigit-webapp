package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/wigit-ng/webapp/backend/internal/api/v1"
	"github.com/wigit-ng/webapp/backend/internal/api/v1/middlewares"
	"github.com/wigit-ng/webapp/backend/internal/config"
	"github.com/wigit-ng/webapp/backend/internal/flags"
)

var (
	// AppConfig serves as the config for the application. Values are gotten
	// from the environment variables.
	AppConfig config.Config

	// GinRouter is the router for the Gin server. It is passed to the standard
	// net/http router to implement graceful shutdown.
	GinRouter *gin.Engine

	// HTTPRouter runs the backend server.
	HTTPRouter *http.Server

	// logOutputFile is the file pointer to the log file.
	logOutputFile *os.File
)

func init() {
	var environment string
	environment, logOutputFile = flags.Parse()

	switch environment {
	case "prod":
		if logOutputFile == nil {
			log.Panic().Msg("failed to create log file for production environment")
		}
	default:
		if err := godotenv.Load(); err != nil {
			log.Panic().Err(err).Msg("failed to load .env file in development environment")
		}

	}

	AppConfig = config.NewConfig(environment)
	if err := middlewares.NewRedis(AppConfig); err != nil {
		log.Panic().Err(err).Msg("failed to initialize redis")
	}
	GinRouter = server.SetAPIRouter(AppConfig)
	HTTPRouter = server.SetWebRouter(GinRouter, AppConfig)
}

// main Entry point to the program
//
//	@contact.name	API Support
//	@contact.url	/contact
//	@contact.email	ecokeke21@gmail.com
func main() {
	server.ListenAndServe(HTTPRouter)
	if logOutputFile != nil {
		defer logOutputFile.Close()
	}
}
