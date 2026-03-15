package transcoding

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("generating test video...")

	videoDirectory := "tmp"
	options := TestVideoOpts{
		FfmpegPath:      "/usr/bin/ffmpeg",
		Resolution:      RES_1080p,
		Framerate:       60,
		DurationSeconds: "10",
		PixelFormat:     DEFAULT_PIXEL_FMT,
		OutFile:         filepath.Join(videoDirectory, "test_video_1080p.mp4"),
	}
	_, err := GenerateTestVideo(options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate test video: %s\n", err)
		os.Exit(1)
	}

	code := m.Run()

	if err := os.RemoveAll(videoDirectory); err != nil {
		fmt.Fprintf(os.Stderr, "failed to clean up generated video: %s\n", err)
		os.Exit(1)
	}

	os.Exit(code)
}
