package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cenkalti/backoff"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/wigit-ng/webapp/backend/internal/config"
)

// RedisClient represents a redis client instance for caching.
var RedisClient *redis.Client

// NewRedis creates a new Redis client based on the provided configuration.
// It establishes a connection to the Redis server using the host and port
// information from the provided `config.Config`. If successful, it assigns
// the client to the global `RedisClient` variable and returns no error.
// If there is an error during initialization, it returns an error.
func NewRedis(cnf config.Config) error {
	db, err := strconv.Atoi(cnf.RedisDB)
	if err != nil {
		return err
	}

	operation := func() error {
		client := redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", cnf.RedisHost, cnf.RedisPort),
				Password: cnf.RedisPass,
				DB:       db,
			},
		)

		if client == nil {
			return fmt.Errorf("failed to initialize redis client")
		}

		RedisClient = client
		return nil
	}

	return backoff.Retry(operation, backoff.NewExponentialBackOff())
}

// Redis is a Gin middleware that checks if the requested URL is
// present in the Redis cache. If found in the cache, it retrieves and responds
// with the cached data. If not found, it continues to the next middleware in
// the request chain.
func Redis(ctx *gin.Context) {
	content, err := RedisClient.Get(ctx, ctx.Request.URL.String()).Bytes()
	if err == redis.Nil {
		ctx.Next()
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Determine type of cached item, whether slice of maps or map.
	var body interface{}
	if content[0] == '[' && content[len(content)-1] == ']' {
		body = new([]map[string]interface{})
	} else {
		body = new(map[string]interface{})
	}

	if err := json.Unmarshal(content, body); err != nil {
		// TODO: Handle JSON unmarshal error.
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": body,
	})
}
