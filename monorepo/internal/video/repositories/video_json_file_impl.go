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
)

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
		log.Fatalf("failed to initialize JSOn file: %s\n", err)
	}

	return instance
}

func (r *videoJSONFileRepository) Create(ctx context.Context, video models.NewVideo) (entities.NewVideo, error) {
	var newVideo entities.NewVideo

	videos, err := r.read()
	if err != nil {
		return newVideo, err
	}

	videos = append(videos, models.Video{
		ID:           video.ID,
		Title:        video.Title,
		Status:       video.Status,
		OriginalPath: video.OriginalPath,
	})

	b, err := json.Marshal(videos)
	if err != nil {
		return newVideo, err
	}

	if err := r.write(b); err != nil {
		return newVideo, err
	}

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
			video = entities.ModelToVideo(vid)
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
				videos[i].MasterPlaylistPath.String = *update.MasterPlaylistPath
			}

			if update.ThumbnailPath != nil {
				videos[i].ThumbnailPath.String = *update.ThumbnailPath
			}

			if update.DurationSeconds != nil {
				videos[i].DurationSeconds.Int64 = int64(*update.DurationSeconds)
			}

			if update.Width != nil {
				videos[i].Width.Int64 = int64(*update.Width)
			}

			if update.Height != nil {
				videos[i].Height.Int64 = int64(*update.Height)
			}

			if update.ProcessedAt != nil {
				videos[i].ProcessedAt.Time = *update.ProcessedAt
			}

			video = entities.ModelToVideo(videos[i])
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

	if _, err := os.Stat(r.filepath); err != nil && errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(r.filepath, make([]byte, 0), r.permission); err != nil {
			return err
		}
	}

	return nil
}

func (r *videoJSONFileRepository) read() ([]models.Video, error) {
	videos := make([]models.Video, 0)

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
