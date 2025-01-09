package repository

import (
	"errors"
	"strings"

	"github.com/c8763yee/mygo-backend/internal/models"
	"gorm.io/gorm"
)

type SentenceRepository struct {
	db *gorm.DB
}

func NewSentenceRepository(db *gorm.DB) *SentenceRepository {
	return &SentenceRepository{db: db}
}

func (r *SentenceRepository) SearchByText(text string, videoName models.VideoNameEnum, episode string, pagedBy, page int) ([]models.SentenceItem, int64, error) {
	var items []models.SentenceItem
	var resultCount int64

	if text == "" {
		return nil, 0, errors.New("text is required")
	}

	// escape `%` in text.
	text = strings.ReplaceAll(text, "%", "\\%")

	query := r.db.Model(&models.SentenceItem{}).Where("text LIKE ?", "%"+text+"%")

	// if episode or videoName is provided, filter by them.
	if episode != "" {
		query = query.Where("episode = ?", episode)
	}
	if videoName != "" {
		query = query.Where("video_name = ?", videoName)
	}

	if err := query.Count(&resultCount).Error; err != nil {
		return nil, 0, err
	}

	if err := (query.Order("segment_id ASC").
		Offset((page - 1) * pagedBy).
		Limit(pagedBy).
		Find(&items).Error); err != nil {
		return nil, 0, err
	}

	return items, resultCount, nil
}
