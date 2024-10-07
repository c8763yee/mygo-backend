package data

import (
	"encoding/json"
	"os"

	"gorm.io/gorm"
)

// SentenceItem represents a sentence item with GORM model tags
type SentenceItem struct {
	ID         uint   `gorm:"type:int;primaryKey;autoIncrement;not null;"`
	Text       string `json:"text" gorm:"type:text;not null;"`
	Episode    string `json:"episode" gorm:"type:varchar(3);not null;"`
	FrameStart uint   `json:"frame_start" gorm:"type:int;not null;"`
	FrameEnd   uint   `json:"frame_end" gorm:"type:int;not null;"`
	SegmentId  uint   `json:"segment_id" gorm:"type:int;not null;index;unique;"`
}

// TableName sets the insert table name for this struct type
func (SentenceItem) TableName() string {
	return "sentence"
}

type Sentence struct {
	gorm.Model
	Result []SentenceItem `json:"result"`
}

func GetDataFromFile() Sentence {
	text, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	sentenceData := Sentence{}
	err = json.Unmarshal(text, &sentenceData)
	if err != nil {
		panic(err)
	}
	return sentenceData
}
