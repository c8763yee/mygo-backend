package handlers

import (
	"github.com/c8763yee/mygo-backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	var req service.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, count, err := service.SearchService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
		"count":   count,
	})
}
