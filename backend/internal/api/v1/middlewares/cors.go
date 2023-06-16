package middlewares

import "github.com/gin-contrib/cors"

// CorsConfig sets up cors configuration based on the slices provided.
// It returns the configuration object.
func CorsConfig(allowOrigins, allowMethods, allowHeaders []string) cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = allowOrigins
	config.AllowMethods = allowMethods
	config.AllowHeaders = allowHeaders
	config.AllowCredentials = true

	return config
}
