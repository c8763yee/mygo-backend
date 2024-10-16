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
		api.POST("/search", middleware.RateLimit(memory, middleware.SearchRateLimit), handlers.Search(db))
		api.GET("/gif", middleware.RateLimit(memory, middleware.GIFRateLimit), handlers.ExtractGIFAsFile())
		api.POST("/extract_gif", middleware.RateLimit(memory, middleware.GIFRateLimit), handlers.ExtractGIF())
		api.GET("/frame", middleware.RateLimit(memory, middleware.FrameRateLimit), handlers.ExtractFrameAsFile())
		api.POST("/extract_frame", middleware.RateLimit(memory, middleware.FrameRateLimit), handlers.ExtractFrame())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	return r
}
