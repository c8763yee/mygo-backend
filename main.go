package main

import (
	"mygo/backend"
	_ "mygo/data"

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
