package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/wigit-gh/webapp/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/internal/config"
	"github.com/wigit-gh/webapp/internal/db"
)

// DBConnector servers as a global link to the database.
var DBConnector *db.DB

// ListenAndServer starts up the gin web server to listen on host and port as
// specified in `conf`.
func ListenAndServer(conf config.Config) {
	middlewares.ConfigureJWT([]byte(conf.JWTSecret))

	// Our link to the database
	DBConnector = db.GetConnector(conf)

	// A configured router with logger and recovery middleware
	router := gin.Default()

	// We don't care about trailing slashes. /example and /example/ mean the same thing
	router.RedirectTrailingSlash = true

	// Create the api group
	api := router.Group("/api/v1")

	// Add all the routes needed
	addRoutes(api)

	// 404 handler
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	// Use http server to implement graceful shutdown
	r := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", conf.GinHost, conf.GinPort),
	}

	// Start the server in a go routine so the graceful shutdown mechanism below can
	// be reached
	go func() {
		if err := r.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Panic().Err(err).Msg("failed to start server")
		}
	}()

	// Create a channel
	quit := make(chan os.Signal)
	// Pass channel to signal to notify when any of the signals are encountered
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Execution will be blocked here until a signal is read from quit
	<-quit
	log.Info().Msg("Attempting to shutdown server gracefully...")

	// Create a timeout of 5 seconds to give the server time to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown gracefully")
	}

	log.Info().Msg("Server gracefully shutdown.")
}

// addRoutes adds all routes needed by the front end app.
func addRoutes(api *gin.RouterGroup) {}
