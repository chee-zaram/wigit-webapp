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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	docs "github.com/wigit-gh/webapp/backend/docs"
	"github.com/wigit-gh/webapp/backend/internal/api/v1/middlewares"
	"github.com/wigit-gh/webapp/backend/internal/api/v1/routes"
	"github.com/wigit-gh/webapp/backend/internal/config"
	"github.com/wigit-gh/webapp/backend/internal/db"
)

// SetAPIRouter configures the gin router with all necessary routes and middleware.
func SetAPIRouter(conf config.Config) *gin.Engine {
	middlewares.CreateSigner([]byte(conf.JWTSecret))
	middlewares.CreateVerifier([]byte(conf.JWTSecret))

	// Our link to the database
	db.Connector = db.GetConnector(conf)

	// A configured router with logger and recovery middleware
	router := gin.Default()

	// Get cors config object.
	corsConfig := middlewares.CorsConfig(
		[]string{"*"},
		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		[]string{"Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept"},
	)

	// Use the cors middleware with our configuration.
	router.Use(cors.New(corsConfig))

	// We don't care about trailing slashes. /example and /example/ mean the same thing
	router.RedirectTrailingSlash = true

	docs.SwaggerInfo.Title = "Wigit Web Application Backend Server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is the backend server for the application."
	docs.SwaggerInfo.Host = "cheezaram.tech"
	// docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}
	// docs.SwaggerInfo.Schemes = []string{"http"}

	// Create the api group
	api := router.Group("/api/v1")

	// Add all the routes needed
	addRoutes(api)

	// Create the admin group and add auth middlewares
	admin := api.Group(
		"/admin",
		middlewares.JWTAuthentication,
		middlewares.AdminAuthorization,
	)

	// Add all admin routes
	addAdminRoutes(admin)

	// Create the super admin group and add auth middlewares
	superAdmin := api.Group(
		"/super_admin",
		middlewares.JWTAuthentication,
		middlewares.SuperAdminAuthorization,
	)
	addSuperAdminRoutes(superAdmin)

	// Add only authentication middleware for all other users
	customer := api.Group("/",
		middlewares.JWTAuthentication,
	)
	addCustomerRoutes(customer)

	// Specify route for swagger
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 404 handler
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

// SetWebRouter uses the gin router to configure a traditional http router.
// This is to implement graceful shutdown.
func SetWebRouter(router *gin.Engine, conf config.Config) *http.Server {
	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", conf.GinHost, conf.GinPort),
	}
}

// ListenAndServe starts up the web server implementing graceful shutdown.
func ListenAndServe(r *http.Server) {
	// Start the server in go routine so graceful shutdown mechanism below can be reached
	go func() {
		if err := r.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Info().Msg("Server now closed")
		} else if err != nil {
			log.Error().Err(err).Msg("Something went wrong while listening")
		}
	}()

	// Create a channel
	quit := make(chan os.Signal, 1)
	// Pass channel to signal to notify when any of the signals are encountered
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// Execution will be blocked here until a signal is read from quit
	<-quit
	log.Info().Msg("Attempting to shutdown server gracefully...")

	// Create a timeout of 5 seconds to give the server time to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown gracefully")
	}

	// Wait for the 5 second timeout
	select {
	case <-ctx.Done():
		log.Info().Msg("Timeout of 5 seconds elapsed")
	}

	log.Info().Msg("Server gracefully shutdown.")
}

// addRoutes adds all routes needed by the front end app.
func addRoutes(api *gin.RouterGroup) {
	routes.SignUpRoutes(api)
	routes.SignInRoutes(api)
	routes.ProductsRoutes(api)
	routes.ServicesRoutes(api)
	routes.SlotsRoutes(api)
	routes.ResetPasswordRoutes(api)
}

// addCustomerRoutes adds all routes that need authentication only.
func addCustomerRoutes(customer *gin.RouterGroup) {
	routes.CartRoutes(customer)
	routes.OrdersRoutes(customer)
	routes.BookingsRoutes(customer)
	routes.UsersRoutes(customer)
}

// addAdminRoutes adds all routes that need both authentication and authorization.
// Typically, this is all admin routes.
func addAdminRoutes(admin *gin.RouterGroup) {
	routes.AdminBookingsRoutes(admin)
	routes.AdminProductsRoutes(admin)
	routes.AdminServicesRoutes(admin)
	routes.AdminSlotsRoutes(admin)
	routes.AdminOrdersRoutes(admin)
	routes.AdminUsersRoutes(admin)
}

// addSuperAdminRoutes adds all routes that need super admin priviledges.
func addSuperAdminRoutes(superAdmin *gin.RouterGroup) {
	routes.SuperAdminUsersRoutes(superAdmin)
}
