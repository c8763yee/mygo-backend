package backend

import (
	"encoding/base64"
	"fmt"
	"mygo/data"
	"mygo/video"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

const (
	pagedBy   = 20
	videoPath = "%s/mygo-anime/%s.mp4" // $HOME/mygo-anime/%$episode.mp4
)

var homePath = os.Getenv("HOME")

func Ping(c *gin.Context) {
	arg := c.Query("arg")
	c.JSON(http.StatusOK, gin.H{
		"message": arg,
	})
}

func Search(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.PagedBy == 0 {
		req.PagedBy = pagedBy
	}

	result, count, err := data.SearchByText(req.Query, req.Episode, req.PagedBy, req.NthPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, SearchResponse{Results: result, Count: int(count)})

}

func ExtractFrame(c *gin.Context) {
	var req ExtraceFrameRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if runtime.GOOS == "windows" {
		homePath = os.Getenv("USERPROFILE")
	}
	videoPath := fmt.Sprintf(videoPath, homePath, req.Episode)
	frame, fps := video.FetchVideoFPS(videoPath)
	if frame < req.FrameNumber {
		c.JSON(http.StatusBadRequest, gin.H{"error": "frame number out of range"})
		return
	}
	buf, err := video.ExtractFrame(req.Episode, req.FrameNumber, fps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ExtraceFrameResponse{Frame: base64.StdEncoding.EncodeToString(buf.Bytes())})
}

func ExtractGIF(c *gin.Context) {
	var req ExtraceGIFRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if runtime.GOOS == "windows" {
		homePath = os.Getenv("USERPROFILE")
	}
	videoPath := fmt.Sprintf(videoPath, homePath, req.Episode)
	frame, fps := video.FetchVideoFPS(videoPath)
	if req.Start < 0 || req.End < 0 || req.End > frame {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start or end frame"})
		return
	}
	buf, err := video.ExtractGIF(req.Episode, req.Start, req.End, fps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ExtraceGIFResponse{GIF: base64.StdEncoding.EncodeToString(buf.Bytes())})
}
