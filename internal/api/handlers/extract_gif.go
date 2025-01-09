package handlers

import (
	"net/http"
	"strconv"

	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// ExtractGIF godoc
//
//	@Summary		Extract GIF
//	@Description	Extract GIF as File based on episode, start, and end
//	@Tags			extract
//	@Accept			json
//	@Produce		image/gif
//	@Param			video_name	query	string	true	"Video Name"
//	@Param			episode		query	string	true	"Episode"
//	@Param			start		query	int		true	"Start"
//	@Param			end			query	int		true	"End"
//	@Success		200			{file}	image/gif
//	@Router			/gif [get]
func ExtractGIF() gin.HandlerFunc {
	return func(c *gin.Context) {
		videoName := c.Query("video_name")
		episode := c.Query("episode")
		start, _ := strconv.Atoi(c.Query("start"))
		end, _ := strconv.Atoi(c.Query("end"))

		videoService := service.NewVideoService()
		buf, err := videoService.ExtractGIF(videoName, episode, start, end)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "image/gif", buf.Bytes())
	}
}
