// middleware/cache.go
package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/c8763yee/mygo-backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var (
	// Define cache durations as constants
	// CacheDuration   = 5 * time.Minute
	// CleanupInterval = 10 * time.Minute
	CacheDuration   = time.Duration(config.AppConfig.Cache.CacheDuration) * time.Second
	CleanupInterval = time.Duration(config.AppConfig.Cache.CleanupInterval) * time.Second
	Cache           = cache.New(CacheDuration, CleanupInterval)
)

func CacheMiddlewareGIF() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.String() // Use full URL as cache key
		fmt.Printf("Cache Duration: %d\n", config.AppConfig.Cache.CacheDuration)
		c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", config.AppConfig.Cache.CacheDuration))
		// Try to get the cached response
		if cached, found := Cache.Get(key); found {
			c.Data(http.StatusOK, "image/gif", cached.([]byte))

			c.Abort()
			return
		}

		// Store the response
		c.Writer = &cacheWriter{ResponseWriter: c.Writer, key: key}
		c.Next()
	}
}

func CacheMiddlewareFrame() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.String() // Use full URL as cache key
		fmt.Printf("Cache Duration: %d\n", config.AppConfig.Cache.CacheDuration)

		c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", config.AppConfig.Cache.CacheDuration))

		// Try to get the cached response
		if cached, found := Cache.Get(key); found {
			c.Data(http.StatusOK, "image/jpeg", cached.([]byte))

			c.Abort()
			return
		}

		// Store the response
		c.Writer = &cacheWriter{ResponseWriter: c.Writer, key: key}
		c.Next()
	}
}

type cacheWriter struct {
	gin.ResponseWriter
	key  string
	body []byte
}

func (w *cacheWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	Cache.Set(w.key, w.body, cache.DefaultExpiration)
	return w.ResponseWriter.Write(b)
}
