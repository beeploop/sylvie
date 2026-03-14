package transcoding

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	DEFAULT_PIXEL_FMT = "yuv420p"
)

type TestVideoOpts struct {
	FfmpegPath      string
	Resolution      Resolution
	Framerate       int
	DurationSeconds string
	PixelFormat     string
	OutFile         string
}

func GenerateTestVideo(input TestVideoOpts) (string, error) {
	dirpath := filepath.Dir(input.OutFile)
	if err := os.MkdirAll(dirpath, os.FileMode(0777)); err != nil {
		return "", err
	}

	cmd := exec.Command(
		input.FfmpegPath,
		"-y",
		"-f", "lavfi",
		"-i", fmt.Sprintf("testsrc=size=%s:rate=%d", input.Resolution.Dimension(), input.Framerate),
		"-f", "lavfi", "-i", "sine=frequency=1000",
		"-t", input.DurationSeconds,
		"-c:v", "libx264",
		"-c:a", "aac",
		"-pix_fmt", input.PixelFormat,
		input.OutFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return input.OutFile, nil
}
