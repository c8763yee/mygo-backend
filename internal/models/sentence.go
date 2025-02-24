package models

type SentenceItem struct {
	ID         uint   `gorm:"type:int;primaryKey;autoIncrement;not null;"`
	Text       string `json:"text" gorm:"type:text;not null;" csv:"text"`
	Episode    string `json:"episode" gorm:"type:varchar(3);not null;" csv:"episode"`
	FrameStart uint   `json:"frame_start" gorm:"type:int;not null;" csv:"frame_start"`
	FrameEnd   uint   `json:"frame_end" gorm:"type:int;not null;" csv:"frame_end"`
	SegmentId  uint   `json:"segment_id" gorm:"type:int;not null;index;unique;" csv:"segment_id"`
	VideoName  string `json:"video_name" gorm:"type:varchar(255);not null;" csv:"video_name"`
}

func (SentenceItem) TableName() string {
	return "sentence"
}

type VideoNameEnum string

const (
	AveMujica VideoNameEnum = "Ave Mujica"
	MyGO      VideoNameEnum = "MyGO"
)

type SearchRequest struct {
	Query     string        `json:"query"`
	VideoName VideoNameEnum `json:"video_name" enums:"Ave Mujica,MyGO"`
	Episode   string        `json:"episode"`
	NthPage   int           `json:"nth_page"`
	PagedBy   int           `json:"paged_by"`
}

type SearchResponse struct {
	Count   int            `json:"count"`
	Results []SentenceItem `json:"results"`
}

type ExtractFrameRequest struct {
	// Available Video Name:
	// * Ave Mujica - Ave Mujica
	// * MyGO - MyGO
	VideoName   VideoNameEnum `json:"video_name" enums:"Ave Mujica,MyGO"`
	Episode     string        `json:"episode"`
	FrameNumber int           `json:"frame"`
}

type ExtractGIFRequest struct {
	// Available Video Name:
	// * Ave Mujica - Ave Mujica
	// * MyGO - MyGO
	VideoName VideoNameEnum `json:"video_name" enums:"Ave Mujica,MyGO"`
	Episode   string        `json:"episode"`
	Start     int           `json:"start"`
	End       int           `json:"end"`
	Format    string        `json:"format" enums:"gif,webm"`
}
