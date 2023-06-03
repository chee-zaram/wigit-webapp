package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/wigit-gh/webapp/internal/api/v1"
	"github.com/wigit-gh/webapp/internal/config"
	"github.com/wigit-gh/webapp/internal/flags"
)

// main Entry point to the program
//
//	@contact.name	API Support
//	@contact.url	contact
//	@contact.email	ecokeke21@gmail.com
func main() {
	env := flags.Parse()
	// NOTE: in production we will specify the environment variables in our service file
	// TODO: configure logger path to use standard /var/log dir.
	/* if env != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Panic().Err(err).Msg("failed to load .env file")
		}
	} */

	if err := godotenv.Load(); err != nil {
		log.Panic().Err(err).Msg("failed to load .env file")
	}

	conf := config.NewConfig(env)
	ginRouter := server.SetGINRouterV1(conf)
	httpRouter := server.SetHTTPRouter(ginRouter, conf)
	server.ListenAndServer(httpRouter)
}
