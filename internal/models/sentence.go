package models

import (
	"gorm.io/gorm"
)

type SentenceItem struct {
	gorm.Model
	Text       string `json:"text" gorm:"type:text;not null;"`
	Episode    string `json:"episode" gorm:"type:varchar(3);not null;"`
	FrameStart uint   `json:"frame_start" gorm:"type:int;not null;"`
	FrameEnd   uint   `json:"frame_end" gorm:"type:int;not null;"`
	SegmentId  uint   `json:"segment_id" gorm:"type:int;not null;index;unique;"`
}

func (SentenceItem) TableName() string {
	return "sentences"
}
