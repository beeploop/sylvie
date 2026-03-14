package transcoding

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranscoder(t *testing.T) {
	ffmpegPath := "/usr/bin/ffmpeg"

	t.Run("test transcoder command builder", func(t *testing.T) {
		outDir := t.TempDir()

		tests := []struct {
			Name              string
			BaseDirPermission os.FileMode
			Input             Rendetion
			ExpectedCmd       []string
		}{
			{
				Name:              "test happy path",
				BaseDirPermission: os.FileMode(0777),
				Input: Rendetion{
					VideoID:    "1234",
					InputPath:  "path/to/video.mp4",
					Resolution: RES_1080p,
				},
				ExpectedCmd: []string{
					ffmpegPath, "-y",
					"-i", "path/to/video.mp4",
					"-vf", "scale=1920:1080",
					"-c:v", "libx264", "-b:v", "5000k",
					"-c:a", "aac", "-b:a", "128k",
					"-f", "hls",
					"-hls_time", "5",
					"-hls_playlist_type", "vod",
					"-hls_segment_filename", outDir + "/1234/1080p/segment_%03d.ts",
					outDir + "/1234/1080p/index.m3u8",
				},
			},
		}

		for _, tc := range tests {
			transcoder := NewTranscoder(ffmpegPath, outDir, tc.BaseDirPermission)

			t.Run(tc.Name, func(t *testing.T) {
				outDir := transcoder.outputDirectory(tc.Input.VideoID, tc.Input.Resolution.Name())
				cmd := transcoder.buildCommand(tc.Input, outDir)

				assert.EqualValues(t, tc.ExpectedCmd, cmd.Args)
			})
		}
	})

	t.Run("test transcoder implementation", func(t *testing.T) {
		testInputDir := "transcoder_test_input"
		outDir := t.TempDir()

		wd, err := os.Getwd()
		if err != nil {
			t.Fatalf("cannot read current working directory: %s\n", err)
		}

		inputDirInfo, err := os.Stat(filepath.Join(wd, testInputDir))
		if errors.Is(err, os.ErrNotExist) {
			t.Fatalf("%s does not exists", testInputDir)
		}
		if err == nil && !inputDirInfo.IsDir() {
			t.Fatalf("%s is not a directory", testInputDir)
		}

		entries, err := os.ReadDir(filepath.Join(wd, testInputDir))
		if err != nil {
			t.Fatalf("cannot read contents of %s directory: %s\n", testInputDir, err)
		}

		tests := []struct {
			Name    string
			VideoID string
			Input   string
		}{}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			tests = append(tests, struct {
				Name    string
				VideoID string
				Input   string
			}{
				Name:    "transcoding integration test",
				VideoID: strings.Split(entry.Name(), ".")[0],
				Input:   filepath.Join(testInputDir, entry.Name()),
			})
		}

		transcoder := NewTranscoder(ffmpegPath, outDir, os.FileMode(0777))

		for _, tc := range tests {
			r := Rendetion{
				VideoID:    tc.VideoID,
				InputPath:  tc.Input,
				Resolution: RES_1080p,
			}

			path, err := transcoder.Transcode(r)
			assert.NoError(t, err)
			assert.NotEmpty(t, path)

			_, err = os.Stat(path)
			assert.NoError(t, err)
		}
	})
}
