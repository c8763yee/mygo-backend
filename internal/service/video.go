package service

import (
	"bytes"
	"errors"

	"github.com/c8763yee/mygo-backend/internal/utils"
)

type VideoService struct{}

const MAX_GIF_FRAMES_DIFF = 1000

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (s *VideoService) ExtractFrame(episode string, frameNumber int) (*bytes.Buffer, error) {
	return utils.ExtractFrame(episode, frameNumber)
}

func (s *VideoService) ExtractGIF(episode string, startFrame, endFrame int) (*bytes.Buffer, error) {
	// raise error if diff between start and end frames is too large (absolutely arbitrary)
	frameDiff := endFrame - startFrame
	if frameDiff < 0 {
		frameDiff = -frameDiff
	}

	if frameDiff > MAX_GIF_FRAMES_DIFF {
		return nil, errors.New("frame diff too large")
	}

	return utils.ExtractGIF(episode, startFrame, endFrame)
}
