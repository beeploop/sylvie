package transcoding

import (
	"errors"
	"strconv"
	"strings"
	"sylvie/internal/video/entities"
)

type ProbeStream struct {
	CodecName   string `json:"codec_name"`
	CodecType   string `json:"codec_type"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FrameRate   string `json:"r_frame_rate"`
	AspectRatio string `json:"display_aspect_ratio"`
}

type ProbeFormat struct {
	Filename   string `json:"filename"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	FormatName string `json:"format_name"`
}

type ProbeResult struct {
	Streams []ProbeStream `json:"streams"`
	Format  ProbeFormat   `json:"format"`
}

func normalizeProbeResult(result ProbeResult) (entities.VideoMetadata, error) {
	size, err := strconv.Atoi(result.Format.Size)
	if err != nil && !errors.Is(err, strconv.ErrSyntax) {
		return entities.VideoMetadata{}, err
	}

	duration, err := strconv.ParseFloat(result.Format.Duration, 64)
	if err != nil {
		return entities.VideoMetadata{}, err
	}

	fps, err := parseFPS(result.Streams[0].FrameRate)
	if err != nil {
		return entities.VideoMetadata{}, err
	}

	metadata := entities.VideoMetadata{
		Width:     result.Streams[0].Width,
		Height:    result.Streams[0].Height,
		Framerate: fps,
		Duration:  duration,
		Size:      size,
		Codec:     result.Streams[0].CodecName,
	}

	return metadata, nil
}

func parseFPS(fps string) (float64, error) {
	parts := strings.Split(fps, "/")

	num, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, err
	}

	den, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, err
	}

	return num / den, nil
}
