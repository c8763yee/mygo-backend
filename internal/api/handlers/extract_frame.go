package handlers

import (
	"fmt"
	"net/http"

	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// ExtractFrame godoc
//
//	@Summary		Extract Frame as File
//	@Description	Extract Frame as File based on episode and frame number
//	@Tags			extract
//	@Accept			json
//	@Produce		image/jpeg
//	@Param			video_name	query	string	true	"Video Name"
//	@Param			episode		query	string	true	"Episode"	enum(1,2,3,4,5,6,7,8,9,10,11,12,13)	default(1)
//	@Param			frame		query	int		true	"Frame Number"
//	@Success		200			{file}	image/jpeg
//	@Router			/frame [get]
func ExtractFrame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var videoName string = c.Query("video_name")
		var episode string = c.Query("episode")
		var frameNumber int

		if _, err := fmt.Sscanf(c.Query("frame"), "%d", &frameNumber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		videoService := service.NewVideoService()
		buf, err := videoService.ExtractFrame(videoName, episode, frameNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Data(http.StatusOK, "image/jpeg", buf.Bytes())
	}
}
