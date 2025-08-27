package utils

import (
	"fmt"
	"time"
)

func SecondsToTimestamp(seconds int) string {
	h := seconds / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func TimestampToSeconds(timestamp string) (int, error) {
	parsed, err := time.Parse("15:04:05", timestamp)
	if err != nil {
		return 0, err
	}
	return parsed.Hour()*3600 + parsed.Minute()*60 + parsed.Second(), nil
}
