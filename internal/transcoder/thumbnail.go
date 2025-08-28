package transcoder

import (
	"os/exec"

	"github.com/beeploop/sylvie/internal/metadata"
	"github.com/beeploop/sylvie/internal/utils"
)

func createDefaultThumbnail(src, dest string, seektimeInSeconds int) (metadata.Thumbnail, error) {
	timestampString := utils.SecondsToTimestamp(seektimeInSeconds)
	timestampInt, err := utils.TimestampToSeconds(timestampString)
	if err != nil {
		return metadata.Thumbnail{}, err
	}

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i",
		src,
		"-ss",
		timestampString,
		"-vframes",
		"1",
		"-q:v",
		"2",
		dest,
	)

	if err := cmd.Run(); err != nil {
		return metadata.Thumbnail{}, err
	}

	return metadata.Thumbnail{
		Timestamp: timestampInt,
		Type:      "default",
		Path:      dest,
	}, nil
}
