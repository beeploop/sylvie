package repositories

import (
	"context"
	"os"
	"path/filepath"
	"sylvie/internal/video/models"
	"testing"
	"time"

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
			Expected []VideoJSON
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
				Expected: []VideoJSON{
					{
						ID:           "1234",
						Title:        "test video",
						OriginalPath: "path/to/video.mp4",
						Status:       string(models.STATUS_UPLOADED),
						CreatedAt:    time.Now(),
					},
				},
			},
			{
				Name:     "test insert empty value",
				Input:    []models.NewVideo{},
				Expected: []VideoJSON{},
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
				assert.Equal(t, len(tc.Expected), len(videos))

				for i := range videos {
					assert.Equal(t, tc.Expected[i].ID, videos[i].ID)
					assert.Equal(t, tc.Expected[i].Title, videos[i].Title)
					assert.Equal(t, tc.Expected[i].OriginalPath, videos[i].OriginalPath)
					assert.Equal(t, tc.Expected[i].Status, videos[i].Status)
					assert.Equal(t, tc.Expected[i].MasterPlaylistPath, videos[i].MasterPlaylistPath)
					assert.Equal(t, tc.Expected[i].Width, videos[i].Width)
					assert.Equal(t, tc.Expected[i].Height, videos[i].Height)
					assert.Equal(t, tc.Expected[i].ThumbnailPath, videos[i].ThumbnailPath)
					assert.Equal(t, tc.Expected[i].DurationSeconds, videos[i].DurationSeconds)
					assert.WithinDuration(t, tc.Expected[i].CreatedAt, videos[i].CreatedAt, time.Second)
				}
			})
		}
	})
}
