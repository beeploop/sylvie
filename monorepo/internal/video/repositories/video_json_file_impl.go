package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sylvie/internal/video/entities"
	"sylvie/internal/video/models"
	"time"
)

type VideoJSON struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`

	OriginalPath       string `json:"original_path"`
	MasterPlaylistPath string `json:"master_playlist_path"`
	ThumbnailPath      string `json:"thumbnail_path"`

	DurationSeconds float64 `json:"duration_seconds"`
	Width           int64   `json:"width"`
	Height          int64   `json:"height"`

	CreatedAt   time.Time `json:"created_at"`
	ProcessedAt time.Time `json:"processed_at"`
}

type videoJSONFileRepository struct {
	filepath   string
	permission os.FileMode
}

func NewVideoJSONFileRepository(filepath string, permission os.FileMode) *videoJSONFileRepository {
	instance := &videoJSONFileRepository{
		filepath:   filepath,
		permission: permission,
	}

	if err := instance.initialize(); err != nil {
		log.Fatalf("failed to initialize JSON file: %s\n", err)
	}

	return instance
}

func (r *videoJSONFileRepository) Create(ctx context.Context, video models.NewVideo) (entities.NewVideo, error) {
	var newVideo entities.NewVideo

	videos, err := r.read()
	if err != nil {
		return newVideo, err
	}

	videos = append(videos, VideoJSON{
		ID:           video.ID,
		Title:        video.Title,
		Status:       string(video.Status),
		OriginalPath: video.OriginalPath,
		CreatedAt:    time.Now(),
	})

	b, err := json.Marshal(videos)
	if err != nil {
		return newVideo, err
	}

	if err := r.write(b); err != nil {
		return newVideo, err
	}

	newVideo.ID = video.ID
	newVideo.Title = video.Title
	newVideo.Status = string(video.Status)
	newVideo.OriginalPath = video.OriginalPath

	return newVideo, nil
}

func (r *videoJSONFileRepository) FindByID(ctx context.Context, id string) (entities.Video, error) {
	var video entities.Video

	videos, err := r.read()
	if err != nil {
		return video, err
	}

	for _, vid := range videos {
		if vid.ID == id {
			video = entities.Video{
				ID:                 vid.ID,
				Title:              vid.Title,
				Status:             vid.Status,
				OriginalPath:       vid.OriginalPath,
				MasterPlaylistPath: vid.MasterPlaylistPath,
				ThumbnailPath:      vid.ThumbnailPath,
				DurationSeconds:    vid.DurationSeconds,
				Width:              int(vid.Width),
				Height:             int(vid.Height),
			}
			return video, nil
		}
	}

	return video, errors.New("video not found")
}

func (r *videoJSONFileRepository) Update(ctx context.Context, id string, update models.UpdateVideo) (entities.Video, error) {
	var video entities.Video

	videos, err := r.read()
	if err != nil {
		return video, err
	}

	for i := range videos {
		if videos[i].ID == id {
			if update.MasterPlaylistPath != nil {
				videos[i].MasterPlaylistPath = *update.MasterPlaylistPath
			}

			if update.ThumbnailPath != nil {
				videos[i].ThumbnailPath = *update.ThumbnailPath
			}

			if update.DurationSeconds != nil {
				videos[i].DurationSeconds = *update.DurationSeconds
			}

			if update.Width != nil {
				videos[i].Width = int64(*update.Width)
			}

			if update.Height != nil {
				videos[i].Height = int64(*update.Height)
			}

			if update.Status != nil {
				videos[i].Status = string(*update.Status)
			}

			if update.ProcessedAt != nil {
				videos[i].ProcessedAt = *update.ProcessedAt
			}

			video = entities.Video{
				ID:                 videos[i].ID,
				Title:              videos[i].Title,
				Status:             videos[i].Status,
				OriginalPath:       videos[i].OriginalPath,
				MasterPlaylistPath: videos[i].MasterPlaylistPath,
				ThumbnailPath:      videos[i].ThumbnailPath,
				DurationSeconds:    videos[i].DurationSeconds,
				Width:              int(videos[i].Width),
				Height:             int(videos[i].Height),
			}
			break
		}
	}

	b, err := json.Marshal(videos)
	if err != nil {
		return video, err
	}

	if err := r.write(b); err != nil {
		return video, err
	}

	return video, nil
}

func (r *videoJSONFileRepository) initialize() error {
	dirpath := filepath.Dir(r.filepath)

	if _, err := os.Stat(dirpath); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(dirpath, r.permission); err != nil {
			return err
		}
	}

	if _, err := os.Stat(r.filepath); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(r.filepath, make([]byte, 0), r.permission); err != nil {
			return err
		}
	}

	return nil
}

func (r *videoJSONFileRepository) read() ([]VideoJSON, error) {
	videos := make([]VideoJSON, 0)

	content, err := os.ReadFile(r.filepath)
	if err != nil {
		return videos, err
	}
	if len(content) == 0 {
		return videos, nil
	}
	if len(content) < 2 {
		return videos, errors.New("malformed json data")
	}

	if err := json.Unmarshal(content, &videos); err != nil {
		return videos, err
	}

	return videos, nil
}

func (r *videoJSONFileRepository) write(data []byte) error {
	return os.WriteFile(r.filepath, data, r.permission)
}
