package service

import (
	"github.com/c8763yee/mygo-backend/internal/models"
	"github.com/c8763yee/mygo-backend/internal/repository"
	"gorm.io/gorm"
)

type SearchService struct {
	repo *repository.SentenceRepository
}

const DEFAULT_PAGED_BY = 20

func NewSearchService(db *gorm.DB) *SearchService {
	return &SearchService{
		repo: repository.NewSentenceRepository(db),
	}
}

func (s *SearchService) SearchByText(text, episode string, pagedBy, page int) ([]models.SentenceItem, int64, error) {
	if pagedBy == 0 {
		pagedBy = DEFAULT_PAGED_BY
	}
	return s.repo.SearchByText(text, episode, pagedBy, page)
}
