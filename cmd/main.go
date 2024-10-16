package main

import (
	"github.com/c8763yee/mygo-backend/docs"
	"github.com/c8763yee/mygo-backend/internal/api"
	"github.com/c8763yee/mygo-backend/internal/config"
	"github.com/c8763yee/mygo-backend/pkg/database"
)

func main() {
	config.LoadConfig()
	r := api.SetupRouter(database.DB)
	docs.SwaggerInfo.BasePath = "/api"
	r.Run(config.AppConfig.ServerAddress)
}

// @title MyGO Backend API
// @version 1.0
// @description This is a server for MyGO Sentence Search and Image/GIF Extraction.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
