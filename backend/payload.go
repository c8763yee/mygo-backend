package backend

import (
	"github.com/c8763yee/mygo-backend/data"

	_ "github.com/gin-gonic/gin"
)

// request of /api/search
type SearchRequest struct {
	Query   string `json:"query"`
	Episode string `json:"episode"`
	NthPage int    `json:"nth_page"`
	PagedBy int    `json:"paged_by"`
}

// response of /api/search
type SearchResponse struct {
	Count   int                 `json:"count"`
	Results []data.SentenceItem `json:"results"`
}

// request of /api/extract_frame
type ExtraceFrameRequest struct {
	Episode     string `json:"episode"`
	FrameNumber int    `json:"frame"`
}

// response of /api/extract_frame
type ExtraceFrameResponse struct {
	Frame string `json:"frame"`
}

// request of /api/extract_gif
type ExtraceGIFRequest struct {
	Episode string `json:"episode"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
}

// response of /api/extract_gif
type ExtraceGIFResponse struct {
	GIF string `json:"gif"`
}
