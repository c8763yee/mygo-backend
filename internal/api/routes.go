package api

import (
	"github.com/c8763yee/mygo-backend/internal/api/handlers"
	"github.com/c8763yee/mygo-backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	memory := memory.NewStore()
	api := r.Group("/api")
	{
		api.POST("/search", middleware.RateLimit(memory, middleware.GetSearchRateLimit()), handlers.Search(db))
		api.GET("/gif", middleware.RateLimit(memory, middleware.GetFrameRateLimit()), middleware.CacheMiddlewareGIF(), handlers.ExtractGIF())
		api.GET("/frame", middleware.RateLimit(memory, middleware.GetGIFRateLimit()), middleware.CacheMiddlewareFrame(), handlers.ExtractFrame())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	return r
}
