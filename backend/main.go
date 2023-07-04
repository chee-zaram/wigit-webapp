package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/wigit-ng/webapp/backend/internal/api/v1"
	"github.com/wigit-ng/webapp/backend/internal/config"
	"github.com/wigit-ng/webapp/backend/internal/flags"
)

// main Entry point to the program
//
//	@contact.name	API Support
//	@contact.url	/contact
//	@contact.email	ecokeke21@gmail.com
func main() {
	environment, logOutputFile := flags.Parse()

	switch environment {
	case "prod":
		if logOutputFile == nil {
			log.Panic().Msg("failed to create log file for production environment")
		}
		defer logOutputFile.Close()
	default:
		if err := godotenv.Load(); err != nil {
			log.Panic().Err(err).Msg("failed to load .env file in development environment")
		}
	}

	appConfig := config.NewConfig(environment)
	ginRouter := server.SetAPIRouter(appConfig)
	httpRouter := server.SetWebRouter(ginRouter, appConfig)
	server.ListenAndServe(httpRouter)
}
