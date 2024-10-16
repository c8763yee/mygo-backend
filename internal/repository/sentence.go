package repository

import (
	"github.com/c8763yee/mygo-backend/internal/models"
	"gorm.io/gorm"
)

type SentenceRepository struct {
	DB *gorm.DB
}

func NewSentenceRepository(db *gorm.DB) *SentenceRepository {
	return &SentenceRepository{DB: db}
}

func (r *SentenceRepository) SearchByText(text, episode string, offset, limit int) ([]models.SentenceItem, int64, error) {
	var items []models.SentenceItem
	var count int64

	query := r.DB.Model(&models.SentenceItem{}).Where("text LIKE ?", "%"+text+"%")
	if episode != "" {
		query = query.Where("episode = ?", episode)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, count, nil
}
