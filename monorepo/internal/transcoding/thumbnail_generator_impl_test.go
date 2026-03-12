package transcoding

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThumbnailGenerator(t *testing.T) {
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
					"ffmpeg",
					"-ss", "3",
					"-i", "/path/to/video.mp4",
					"-frames:v", "1",
					outDir + "/1234/thumbnail.jpg",
				},
			},
		}

		for _, tc := range tests {
			generator := NewThumbnailGenerator(outDir, tc.BaseDirPermission)

			t.Run(tc.Name, func(t *testing.T) {
				outDir := generator.outputDirectory(tc.Input.VideoID)
				cmd := generator.buildCommand(tc.Input.Filepath, outDir)
				assert.EqualValues(t, tc.ExpectedCmd, cmd.Args)
			})
		}
	})

	t.Run("test thumbnail generator implemntation", func(t *testing.T) {
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

		generator := NewThumbnailGenerator(outDir, os.FileMode(0777))

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
