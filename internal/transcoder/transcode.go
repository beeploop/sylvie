package transcoder

import (
	"fmt"
	"os/exec"
	"strconv"
)

func Transcode(inputPath, outputPath string, template Template) error {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i",
		inputPath,
		"-vf",
		fmt.Sprintf("scale=%d:%d", template.Width, template.Height),
		"-c:v",
		template.VideoCodec,
		"-preset",
		template.Preset,
		"-crf",
		strconv.Itoa(template.CRF),
		"-b:v",
		template.VideoBitRate,
		"-c:a",
		template.AudioCodec,
		"-b:a",
		template.AudioBitRate,
		outputPath,
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
