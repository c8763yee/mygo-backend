package main

import (
	"github.com/c8763yee/mygo-backend/internal/api"
	"github.com/c8763yee/mygo-backend/internal/config"
	"github.com/c8763yee/mygo-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	models.InitDB()

	r := gin.Default()
	api.SetupRoutes(r)

	r.Run(config.AppConfig.ServerAddress)
}
