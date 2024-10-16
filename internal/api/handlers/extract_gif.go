package handlers

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/c8763yee/mygo-backend/internal/models"
	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// ExtractGIF godoc
// @Summary Extract GIF
// @Description Extract GIF based on episode, start, and end
// @Tags extract
// @Accept json
// @Produce json
// @Param request body models.ExtractGIFRequest true "Extract GIF parameters"
// @Success 200 {object} models.ExtractGIFResponse
// @Router /extract_gif [post]
func ExtractGIF() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.ExtractGIFRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		videoService := service.NewVideoService()
		buf, err := videoService.ExtractGIF(req.Episode, req.Start, req.End)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.ExtractGIFResponse{GIF: base64.StdEncoding.EncodeToString(buf.Bytes())})
	}
}

// ExtractGIFAsFile godoc
// @Summary Extract GIF as File
// @Description Extract GIF as File based on episode, start, and end
// @Tags extract
// @Accept json
// @Produce image/gif
// @Param episode query string true "Episode"
// @Param start query int true "Start"
// @Param end query int true "End"
// @Success 200 {file} image/gif
// @Router /gif [get]
func ExtractGIFAsFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		episode := c.Query("episode")
		start, _ := strconv.Atoi(c.Query("start"))
		end, _ := strconv.Atoi(c.Query("end"))
		videoService := service.NewVideoService()
		buf, err := videoService.ExtractGIF(episode, start, end)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "image/gif", buf.Bytes())
	}
}
