package models

type SentenceItem struct {
	ID         uint   `gorm:"type:int;primaryKey;autoIncrement;not null;"`
	Text       string `json:"text" gorm:"type:text;not null;"`
	Episode    string `json:"episode" gorm:"type:varchar(3);not null;"`
	FrameStart uint   `json:"frame_start" gorm:"type:int;not null;"`
	FrameEnd   uint   `json:"frame_end" gorm:"type:int;not null;"`
	SegmentId  uint   `json:"segment_id" gorm:"type:int;not null;index;unique;"`
}

func (SentenceItem) TableName() string {
	return "sentence"
}

type SearchRequest struct {
	Query   string `json:"query"`
	Episode string `json:"episode"`
	NthPage int    `json:"nth_page"`
	PagedBy int    `json:"paged_by"`
}

type SearchResponse struct {
	Count   int            `json:"count"`
	Results []SentenceItem `json:"results"`
}

type ExtractFrameRequest struct {
	Episode     string `json:"episode"`
	FrameNumber int    `json:"frame"`
}

type ExtractFrameResponse struct {
	Frame string `json:"frame"`
}

type ExtractGIFRequest struct {
	Episode string `json:"episode"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
}

type ExtractGIFResponse struct {
	GIF string `json:"gif"`
}
