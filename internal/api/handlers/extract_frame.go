package handlers

import (
	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/c8763yee/mygo-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ExtractFrame(c *gin.Context) {
	var req service.ExtractFrameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	frameData, err := service.ExtractFrameService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", frameData)
}

func ExtractFrameAsFile(c *gin.Context) {
	episode := c.Query("episode")
	frameStr := c.Query("frame")
	frame, err := strconv.Atoi(frameStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid frame number"})
		return
	}

	filePath, err := utils.ExtractFrameAsFile(episode, frame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(filePath)
}
