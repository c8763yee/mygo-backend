package main

import (
	"github.com/c8763yee/mygo-backend/backend"
	_ "github.com/c8763yee/mygo-backend/data"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func main() {
	r := backend.CreateEngine()
	res := r.Run("0.0.0.0:8080")
	if res != nil {
		panic(res)
	}
}
