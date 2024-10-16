package service

import (
	"bytes"

	"github.com/c8763yee/mygo-backend/internal/utils"
)

type VideoService struct{}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (s *VideoService) ExtractFrame(episode string, frameNumber int) (*bytes.Buffer, error) {
	return utils.ExtractFrame(episode, frameNumber)
}

func (s *VideoService) ExtractGIF(episode string, startFrame, endFrame int) (*bytes.Buffer, error) {
	return utils.ExtractGIF(episode, startFrame, endFrame)
}
