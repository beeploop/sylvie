package transcoder

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/beeploop/sylvie/internal/metadata"
)

type FfprobeStream struct {
	CodecName   string `json:"codec_name"`
	CodecType   string `json:"codec_type"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FrameRate   string `json:"r_frame_rate"`
	AspectRatio string `json:"display_aspect_ratio"`
}

type FfprobeFormat struct {
	Filename   string `json:"filename"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	FormatName string `json:"format_name"`
}

type FfprobeResult struct {
	Streams []FfprobeStream `json:"streams"`
	Format  FfprobeFormat   `json:"format"`
}

func (r FfprobeResult) primaryVideoSource() (FfprobeStream, bool) {
	for _, stream := range r.Streams {
		if stream.CodecType == "video" {
			return stream, true
		}
	}

	return FfprobeStream{}, false
}

func (r FfprobeResult) primaryAudioSource() (FfprobeStream, bool) {
	for _, stream := range r.Streams {
		if stream.CodecType == "audio" {
			return stream, true
		}
	}

	return FfprobeStream{}, false
}

func (r FfprobeResult) framerate() float32 {
	video, found := r.primaryVideoSource()
	if !found {
		return 0
	}

	parts := strings.Split(video.FrameRate, "/")
	part1, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return 0
	}

	part2, err := strconv.ParseFloat(parts[1], 32)
	if err != nil {
		return 0
	}

	return float32(part1 / part2)
}

func (r FfprobeResult) toSourceMetadata() (metadata.SourceMetadata, error) {
	video, found := r.primaryVideoSource()
	if !found {
		err := fmt.Errorf("primary video stream not found")
		return metadata.SourceMetadata{}, err
	}

	audio, found := r.primaryAudioSource()
	if !found {
		err := fmt.Errorf("primary audio stream not found")
		return metadata.SourceMetadata{}, err
	}

	duration, err := strconv.ParseFloat(r.Format.Duration, 32)
	if err != nil {
		return metadata.SourceMetadata{}, err
	}

	filesize, err := strconv.Atoi(r.Format.Size)
	if err != nil {
		return metadata.SourceMetadata{}, err
	}

	return metadata.SourceMetadata{
		Duration:         float32(duration),
		AspectRatio:      video.AspectRatio,
		VideoCodec:       video.CodecName,
		AudioCodec:       audio.CodecName,
		FrameRate:        r.framerate(),
		OriginalFileSize: filesize,
	}, nil
}

func extractDataWithFfprobe(input string) (FfprobeResult, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v",
		"quiet",
		"-print_format",
		"json",
		"-show_format",
		"-show_streams",
		input,
	)

	jsonData, err := cmd.Output()
	if err != nil {
		return FfprobeResult{}, err
	}

	var result FfprobeResult
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return FfprobeResult{}, err
	}

	return result, nil
}
