package transcoding

import (
	"fmt"
	"strconv"
	"strings"
)

type Resolution string

const (
	RES_1080p Resolution = "1080p"
	RES_720p  Resolution = "720p"
	RES_360p  Resolution = "360p"
)

func (r Resolution) Name() string {
	return string(r)
}

func (r Resolution) VideoBitRate() int {
	switch r {
	case RES_1080p:
		return 5_000_000
	case RES_720p:
		return 2_800_000
	case RES_360p:
		return 800_000
	default:
		return 0
	}
}

func (r Resolution) AudioBitRate() int {
	switch r {
	case RES_1080p:
		return 128_000
	case RES_720p:
		return 128_000
	case RES_360p:
		return 96_000
	default:
		return 0
	}
}

func (r Resolution) Ratio() string {
	switch r {
	case RES_1080p:
		return "1920:1080"
	case RES_720p:
		return "1280:720"
	case RES_360p:
		return "640:360"
	default:
		return "unknown resolution"
	}
}

func (r Resolution) Dimension() string {
	return strings.ReplaceAll(r.Ratio(), ":", "x")
}

func (r Resolution) Width() int {
	width, err := strconv.Atoi(strings.Split(r.Ratio(), ":")[0])
	if err != nil {
		return 0
	}
	return width
}

func (r Resolution) Height() int {
	height, err := strconv.Atoi(strings.Split(r.Ratio(), ":")[1])
	if err != nil {
		return 0
	}
	return height
}

func ResolutionFromName(name string) Resolution {
	switch name {
	case "1080p":
		return RES_1080p
	case "720p":
		return RES_720p
	case "360p":
		return RES_360p
	default:
		return RES_720p
	}
}

func ToBitrateSuffixNotation(bitrate int) string {
	kbps := bitrate / 1_000
	return fmt.Sprintf("%dk", kbps)
}
