package handlers

import (
	"fmt"
	"net/http"

	"github.com/c8763yee/mygo-backend/internal/models"
	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Search godoc
//
//	@Summary		Search for sentences
//	@Description	Search for sentences based on query and other parameters
//	@Tags			search
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SearchRequest	true	"Search parameters"
//	@Success		200		{object}	models.SearchResponse
//	@Router			/search [post]
func Search(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.SearchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		searchService := service.NewSearchService(db)
    fmt.Printf("-------------- [Printf]: Querying %s in %s episode %s (Page %d, paged by %d)----------------\n", req.Query, req.VideoName, req.Episode, req.NthPage, req.PagedBy)
		result, count, err := searchService.SearchByText(req.Query, req.VideoName, req.Episode, req.PagedBy, req.NthPage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.SearchResponse{Results: result, Count: int(count)})
	}
}
