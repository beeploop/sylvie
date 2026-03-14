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
	options := TestVideoOpts{
		FfmpegPath:      "/usr/bin/ffmpeg",
		Resolution:      RES_1080p,
		Framerate:       60,
		DurationSeconds: "10",
		PixelFormat:     DEFAULT_PIXEL_FMT,
		OutFile:         filepath.Join("tmp", "test_video_1080p.mp4"),
	}
	_, err := GenerateTestVideo(options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate test video: %s\n", err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
