package extract

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/c8763yee/mygo-backend/internal/config"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gopkg.in/vansante/go-ffprobe.v2"
)

func FetchVideoFPS(videoPath string) (int, float64) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, videoPath)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}
	frame, err := strconv.Atoi(data.Streams[0].NbFrames)
	if err != nil {
		log.Panicf("Error getting frame: %v", err)
	}

	var num, den float64
	fmt.Sscanf(data.Streams[0].RFrameRate, "%f/%f", &num, &den)
	fps := num / den
	return frame, fps
}

// convert frame to ffmpeg time format("HH:MM:SS.mmm")
func FrameToTime(frame int, fps float64) string {
	sec := float64(frame) / fps
	min := int(sec / 60)
	hour := int(min / 60)
	sec = sec - float64(min*60)
	min = min % 60
	return fmt.Sprintf("%02d:%02d:%06.3f", hour, min, sec)
}

func ExtractFrame(videoName, episode string, frameNumber int) (*bytes.Buffer, error) {
	if frameNumber < 0 {
		return nil, fmt.Errorf("frame number must be positive")
	}

	videoPath := fmt.Sprintf("%s/%s/%s.mp4", config.AppConfig.VideoPath, videoName, episode)
	fmt.Println(videoPath)
	_, fps := FetchVideoFPS(videoPath)

	buf := &bytes.Buffer{}
	err := ffmpeg.Input(videoPath, ffmpeg.KwArgs{"ss": FrameToTime(frameNumber, fps)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg", "loglevel": "quiet"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// func ExtractGIF(episode string, startFrame, endFrame int) (*bytes.Buffer, error) {
func ExtractGIF(videoName, episode string, startFrame, endFrame int) (*bytes.Buffer, error) {
	videoPath := fmt.Sprintf("%s/%s/%s.mp4", config.AppConfig.VideoPath, videoName, episode)
	_, fps := FetchVideoFPS(videoPath)

	reverse := false

	if startFrame > endFrame {
		startFrame, endFrame = endFrame, startFrame
		reverse = true
	} else if startFrame == endFrame {
		return ExtractFrame(videoName, episode, startFrame)
	}

	startTime := FrameToTime(startFrame, fps)
	endTime := FrameToTime(endFrame, fps)

	buf := &bytes.Buffer{}
	input := ffmpeg.Input(videoPath, ffmpeg.KwArgs{"ss": startTime, "to": endTime})

	if reverse {
		input = input.Filter("reverse", nil)
	}

	split := input.Split()
	palette := split.Get("0").Filter("palettegen", nil)
	processPalette := ffmpeg.Filter([]*ffmpeg.Stream{split.Get("1"), palette}, "paletteuse", nil).
		Output("pipe:", ffmpeg.KwArgs{"vcodec": "gif", "format": "gif", "loglevel": "quiet"}).
		WithOutput(buf, os.Stdout)

	fmt.Printf("Extracting GIF from %s: start_frame=%d, end_frame=%d\n", videoPath, startFrame, endFrame)
	err := processPalette.Run()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
