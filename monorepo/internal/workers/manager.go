package workers

import (
	"context"
	"log"
	"os"
	"sylvie/internal/config"
	"sylvie/internal/queue"
	"sylvie/internal/transcoding"
	"sylvie/internal/video/models"
	"sylvie/internal/video/repositories"
	"time"
)

type Manager struct {
	Videos     repositories.VideoRepository
	Probe      transcoding.Probe
	Transcoder transcoding.Transcoder
	Thumbnails transcoding.ThumbnailGenerator
	Playlist   transcoding.PlaylistGenerator
}

func NewManager(config *config.Config) *Manager {
	videoRepository := repositories.NewVideoJSONFileRepository("tmp/db.json", os.FileMode(0777))

	probe := transcoding.NewProbeImpl(config.FFMPEG.FfprobePath)
	transcoder := transcoding.NewTranscoder(config.FFMPEG.FfmpegPath, "tmp/transcoded/", os.FileMode(0777))
	thumbnailGenerator := transcoding.NewThumbnailGenerator(config.FFMPEG.FfmpegPath, "tmp/transcoded/", os.FileMode(0777))
	playlistGenerator := transcoding.NewPlaylistGenerator("tmp/transcoded/", os.FileMode(0777))

	return &Manager{
		Videos:     videoRepository,
		Probe:      probe,
		Transcoder: transcoder,
		Thumbnails: thumbnailGenerator,
		Playlist:   playlistGenerator,
	}
}

func (m *Manager) Handle(job queue.Job) error {
	log.Printf("started processing: %s\n", job.VideoID)

	ctx := context.Background()

	video, err := m.Videos.FindByID(ctx, job.VideoID)
	if err != nil {
		return err
	}

	now := time.Now()
	if _, err := m.Videos.Update(ctx, job.VideoID, models.UpdateVideo{
		Status:      &models.STATUS_PROCESSING,
		ProcessedAt: &now,
	}); err != nil {
		return err
	}

	metadata, err := m.Probe.Analyze(video.OriginalPath)
	if err != nil {
		_, err := m.Videos.Update(ctx, job.VideoID, models.UpdateVideo{
			Status: &models.STATUS_FAILED,
		})
		if err != nil {
			return err
		}
	}

	if _, err := m.Videos.Update(ctx, job.VideoID, models.UpdateVideo{
		DurationSeconds: &metadata.Duration,
		Width:           &metadata.Width,
		Height:          &metadata.Height,
	}); err != nil {
		return err
	}

	return nil
}
