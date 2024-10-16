package api

import (
	"github.com/c8763yee/mygo-backend/internal/api/handlers"
	"github.com/c8763yee/mygo-backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		api.POST("/search", handlers.Search)
		api.POST("/extract_frame", handlers.ExtractFrame)
		api.POST("/extract_gif", handlers.ExtractGIF)
		api.GET("/frame", handlers.ExtractFrameAsFile)
	}
}
