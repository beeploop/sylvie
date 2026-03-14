package repositories

import (
	"context"
	"os"
	"path/filepath"
	"sylvie/internal/video/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoJSONFileRepository(t *testing.T) {
	t.Run("test initialize", func(t *testing.T) {
		outFile := filepath.Join(t.TempDir(), "uploads.json")
		NewVideoJSONFileRepository(outFile, os.FileMode(0777))

		info, err := os.Stat(outFile)
		assert.NoError(t, err)
		assert.NotNil(t, info)
		assert.True(t, info.IsDir() == false)
	})

	t.Run("test write/create", func(t *testing.T) {
		tests := []struct {
			Name     string
			Input    []models.NewVideo
			Expected []models.Video
		}{
			{
				Name: "test happy path",
				Input: []models.NewVideo{
					{
						ID:           "1234",
						Title:        "test video",
						OriginalPath: "path/to/video.mp4",
						Status:       models.STATUS_UPLOADED,
					},
				},
				Expected: []models.Video{
					{
						ID:           "1234",
						Title:        "test video",
						OriginalPath: "path/to/video.mp4",
						Status:       models.STATUS_UPLOADED,
					},
				},
			},
			{
				Name:     "test insert empty value",
				Input:    []models.NewVideo{},
				Expected: []models.Video{},
			},
		}

		for _, tc := range tests {
			outFile := filepath.Join(t.TempDir(), "uploads.json")
			repo := NewVideoJSONFileRepository(outFile, os.FileMode(0777))

			t.Run(tc.Name, func(t *testing.T) {
				for _, input := range tc.Input {
					video, err := repo.Create(context.Background(), input)
					assert.NoError(t, err)
					assert.NotNil(t, video)
				}

				videos, err := repo.read()
				assert.NoError(t, err)
				assert.EqualValues(t, tc.Expected, videos)
			})
		}
	})
}
