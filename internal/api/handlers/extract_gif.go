package handlers

import (
	"net/http"
	"strconv"

	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// ExtractGIF godoc
//
//	@Summary		Extract GIF or WebM
//	@Description	Extract GIF or WebM as File based on episode, start, and end
//	@Tags			extract
//	@Accept			json
//	@Produce		image/gif,video/webm
//	@Param			video_name	query	string	true	"Video Name"
//	@Param			episode		query	string	true	"Episode"
//	@Param			start		query	int		true	"Start"
//	@Param			end			query	int		true	"End"
//	@Param			format		query	string	true	"Format (gif or webm)"
//	@Success		200			{file}	file
//	@Router			/gif [get]
func ExtractGIF() gin.HandlerFunc {
	return func(c *gin.Context) {
		videoName := c.Query("video_name")
		episode := c.Query("episode")
		start, errStart := strconv.Atoi(c.Query("start"))
		end, errEnd := strconv.Atoi(c.Query("end"))
		if errStart != nil || errEnd != nil { // if start or end is not a number, set it to 0
			start = 0
			end = 0
		}
		format := c.Query("format")
		if format != "gif" && format != "webm" { // force to gif if format is not gif or webm
			format = "gif"
		}

		videoService := service.NewVideoService()
		buf, err := videoService.ExtractGIF(videoName, episode, start, end, format)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mimeType := "image/gif"
		if format == "webm" {
			mimeType = "video/webm"
		}
		c.Data(http.StatusOK, mimeType, buf.Bytes())
	}
}
