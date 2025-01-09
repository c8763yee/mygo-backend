package middleware

import (
	"time"

	"github.com/c8763yee/mygo-backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

/*
Rate Limit:
    Search: 15 requests per second/ip
    GIF: 5 requests per second/ip
    Frame: 10 requests per second/ip
*/

var (
	SearchRateLimit *limiter.Rate = &limiter.Rate{
		Period: time.Duration(config.AppConfig.RateLimit.Search.Duration) * time.Second,
		Limit:  config.AppConfig.RateLimit.Search.Limit,
	}

	FrameRateLimit *limiter.Rate = &limiter.Rate{
		Period: time.Duration(config.AppConfig.RateLimit.Frame.Duration) * time.Second,
		Limit:  config.AppConfig.RateLimit.Frame.Limit,
	}

	GIFRateLimit *limiter.Rate = &limiter.Rate{
		Period: time.Duration(config.AppConfig.RateLimit.GIF.Duration) * time.Second,
		Limit:  config.AppConfig.RateLimit.GIF.Limit,
	}
)

func RateLimit(store limiter.Store, rate *limiter.Rate) gin.HandlerFunc {
	return mgin.NewMiddleware(limiter.New(store, *rate))
}
