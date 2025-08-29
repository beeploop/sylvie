package utils

import "fmt"

func BitrateToFfmpegStyleString(bitrate int) string {
	kbps := bitrate / 1_000
	return fmt.Sprintf("%dk", kbps)
}
