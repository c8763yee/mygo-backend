package backend

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateEngine() *gin.Engine {
	var r = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://c8763yee.github.io", "http://localhost:*"}
	r.Use(cors.New(corsConfig))
	r.POST("/api/search", Search)
	r.POST("/api/extract_frame", ExtractFrame)
	r.POST("/api/extract_gif", ExtractGIF)
	r.GET("/api/frame", ExtractFrameAsFile)
	return r
}
