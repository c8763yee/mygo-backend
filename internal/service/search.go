package service

import (
	"github.com/c8763yee/mygo-backend/internal/models"
	"github.com/c8763yee/mygo-backend/internal/repository"
)

type SearchRequest struct {
	Query   string `json:"query"`
	Episode string `json:"episode"`
	NthPage int    `json:"nth_page"`
	PagedBy int    `json:"paged_by"`
}

func SearchService(req SearchRequest) ([]models.SentenceItem, int64, error) {
	if req.PagedBy == 0 {
		req.PagedBy = 20
	}

	offset := (req.NthPage - 1) * req.PagedBy
	repo := repository.NewSentenceRepository(models.DB)
	return repo.SearchByText(req.Query, req.Episode, offset, req.PagedBy)
}
