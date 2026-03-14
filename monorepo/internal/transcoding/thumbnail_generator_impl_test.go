package transcoding

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThumbnailGenerator(t *testing.T) {
	ffmpegPath := "/usr/bin/ffmpeg"

	t.Run("test thumbnail generator build command", func(t *testing.T) {
		outDir := t.TempDir()

		tests := []struct {
			Name              string
			BaseDirPermission os.FileMode
			Input             ThumbnailInput
			ExpectedCmd       []string
		}{
			{
				Name:              "test happy path",
				BaseDirPermission: os.FileMode(0777),
				Input: ThumbnailInput{
					VideoID:  "1234",
					Filepath: "/path/to/video.mp4",
				},
				ExpectedCmd: []string{
					ffmpegPath,
					"-ss", "3",
					"-i", "/path/to/video.mp4",
					"-frames:v", "1",
					outDir + "/1234/thumbnail.jpg",
				},
			},
		}

		for _, tc := range tests {
			generator := NewThumbnailGenerator(ffmpegPath, outDir, tc.BaseDirPermission)

			t.Run(tc.Name, func(t *testing.T) {
				outDir := generator.outputDirectory(tc.Input.VideoID)
				cmd := generator.buildCommand(tc.Input.Filepath, outDir)
				assert.EqualValues(t, tc.ExpectedCmd, cmd.Args)
			})
		}
	})

	t.Run("test thumbnail generation", func(t *testing.T) {
		outDir := t.TempDir()

		tests := []struct {
			Name    string
			VideoID string
			Input   string
		}{
			{
				Name:    "test happy path",
				VideoID: "1234",
				Input:   "tmp/test_video_1080p.mp4",
			},
		}

		generator := NewThumbnailGenerator(ffmpegPath, outDir, os.FileMode(0777))

		for _, tc := range tests {
			t.Run(tc.Name, func(t *testing.T) {
				input := ThumbnailInput{
					VideoID:  tc.VideoID,
					Filepath: tc.Input,
				}

				path, err := generator.Generate(input)
				assert.NoError(t, err)
				assert.NotEmpty(t, path)

				_, err = os.Stat(path)
				assert.NoError(t, err)
			})
		}
	})
}
