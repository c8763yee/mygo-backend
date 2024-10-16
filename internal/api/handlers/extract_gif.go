package handlers

import (
	"net/http"

	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func ExtractGIF(c *gin.Context) {
	var req service.ExtractGIFRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gifData, err := service.ExtractGIFService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/gif", gifData)
}
