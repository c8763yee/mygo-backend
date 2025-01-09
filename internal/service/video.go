package service

import (
	"bytes"
	"errors"

	"github.com/c8763yee/mygo-backend/pkg/extract"
)

type VideoService struct{}

const MAX_GIF_FRAMES_DIFF = 240

func NewVideoService() *VideoService {
	return &VideoService{}
}

// func (s *VideoService) ExtractFrame(episode string, frameNumber int) (*bytes.Buffer, error) {
func (s *VideoService) ExtractFrame(videoName, episode string, frameNumber int) (*bytes.Buffer, error) {
	return extract.ExtractFrame(videoName, episode, frameNumber)
}

func (s *VideoService) ExtractGIF(videoName, episode string, startFrame, endFrame int) (*bytes.Buffer, error) {
	// raise error if diff between start and end frames is too large (absolutely arbitrary)
	frameDiff := endFrame - startFrame
	if frameDiff < 0 {
		frameDiff = -frameDiff
	}

	if frameDiff > MAX_GIF_FRAMES_DIFF {
		return nil, errors.New("frame diff too large")
	}

	return extract.ExtractGIF(videoName, episode, startFrame, endFrame)
}
