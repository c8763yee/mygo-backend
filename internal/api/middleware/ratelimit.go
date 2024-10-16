package middleware

import (
	"time"

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
		Period: 1 * time.Minute,
		Limit:  60,
	}

	FrameRateLimit *limiter.Rate = &limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  30,
	}

	GIFRateLimit *limiter.Rate = &limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  10,
	}
)

func RateLimit(store limiter.Store, rate *limiter.Rate) gin.HandlerFunc {
	return mgin.NewMiddleware(limiter.New(store, *rate))
}
