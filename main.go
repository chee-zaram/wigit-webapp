package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/wigit-gh/webapp/internal/api/v1"
	"github.com/wigit-gh/webapp/internal/config"
	"github.com/wigit-gh/webapp/internal/flags"
)

func main() {
	env := flags.Parse()
	if err := godotenv.Load(); err != nil {
		log.Panic().Err(err).Msg("failed to load .env file")
	}

	conf := config.NewConfig(env)
	server.ListenAndServer(conf)
}
