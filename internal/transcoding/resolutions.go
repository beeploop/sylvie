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
	RES_240p  Resolution = "240p"
	RES_144p  Resolution = "144p"
)

var AllResolutions = []Resolution{
	RES_1080p,
	RES_720p,
	RES_360p,
	RES_240p,
	RES_144p,
}

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
	case RES_240p:
		return 400_000
	case RES_144p:
		return 150_000
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
	case RES_240p:
		return 64_000
	case RES_144p:
		return 48_000
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
	case RES_240p:
		return "426:240"
	case RES_144p:
		return "256:144"
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
	case "240p":
		return RES_240p
	case "144p":
		return RES_144p
	default:
		return RES_144p
	}
}

func ResolutionFromDimension(dimension string) Resolution {
	switch dimension {
	case "1920x1080":
		return RES_1080p
	case "1280x720":
		return RES_720p
	case "640x360":
		return RES_360p
	case "426x240":
		return RES_240p
	case "256x144":
		return RES_144p
	default:
		return RES_144p
	}
}

func ToBitrateSuffixNotation(bitrate int) string {
	kbps := bitrate / 1_000
	return fmt.Sprintf("%dk", kbps)
}

func SelectResolutions(height int) []Resolution {
	resolutions := make([]Resolution, 0)

	for _, res := range AllResolutions {
		if height >= res.Height() {
			resolutions = append(resolutions, res)
		}
	}

	return resolutions
}
