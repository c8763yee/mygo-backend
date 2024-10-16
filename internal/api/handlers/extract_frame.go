package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/c8763yee/mygo-backend/internal/models"
	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// ExtractFrame godoc
// @Summary Extract Frame
// @Description Extract Frame based on episode and frame number
// @Tags extract
// @Accept json
// @Produce json
// @Param request body models.ExtractFrameRequest true "Extract Frame parameters"
// @Success 200 {object} models.ExtractFrameResponse
// @Router /extract_frame [post]
func ExtractFrame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.ExtractFrameRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		videoService := service.NewVideoService()
		buf, err := videoService.ExtractFrame(req.Episode, req.FrameNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.ExtractFrameResponse{Frame: base64.StdEncoding.EncodeToString(buf.Bytes())})
	}
}

// ExtractFrameAsFile godoc
// @Summary Extract Frame as File
// @Description Extract Frame as File based on episode and frame number
// @Tags extract
// @Accept json
// @Produce image/webp
// @Param episode query string true "Episode"
// @Param frame query int true "Frame Number"
// @Success 200 {file} image/webp
// @Router /frame [get]
func ExtractFrameAsFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var episode string = c.Query("episode")
		var frameNumber int
		if _, err := fmt.Sscanf(c.Query("frame"), "%d", &frameNumber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		videoService := service.NewVideoService()
		buf, err := videoService.ExtractFrame(episode, frameNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "image/webp", buf.Bytes())
	}
}
