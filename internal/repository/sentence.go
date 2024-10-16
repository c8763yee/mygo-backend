package repository

import (
	"github.com/c8763yee/mygo-backend/internal/models"
	"gorm.io/gorm"
)

type SentenceRepository struct {
	db *gorm.DB
}

func NewSentenceRepository(db *gorm.DB) *SentenceRepository {
	return &SentenceRepository{db: db}
}

func (r *SentenceRepository) SearchByText(text, episode string, pagedBy, page int) ([]models.SentenceItem, int64, error) {
	var items []models.SentenceItem
	var count int64
	query := r.db.Model(&models.SentenceItem{}).Where("text LIKE ?", "%"+text+"%")

	if episode != "" {
		query = query.Where("episode = ?", episode)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("segment_id ASC").
		Offset((page - 1) * pagedBy).
		Limit(pagedBy).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, count, nil
}
