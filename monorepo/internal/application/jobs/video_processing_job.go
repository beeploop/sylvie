package jobs

import (
	"context"
	"log"
	"sylvie/internal/application"
	"sylvie/internal/queue"
	"sylvie/internal/transcoding"
	"sylvie/internal/video/models"
	"sylvie/internal/video/repositories"
	"time"
)

type VideoProcessingHandler struct {
	Videos     repositories.VideoRepository
	Probe      transcoding.Probe
	Transcoder transcoding.Transcoder
	Thumbnails transcoding.ThumbnailGenerator
	Playlist   transcoding.PlaylistGenerator
}

func NewVideoProcessingHandler(app *application.WorkerApplication) *VideoProcessingHandler {
	return &VideoProcessingHandler{
		Videos:     app.Videos,
		Probe:      app.Probe,
		Transcoder: app.Transcoder,
		Thumbnails: app.Thumbnails,
		Playlist:   app.Playlist,
	}
}

func (h *VideoProcessingHandler) Handle(job queue.Job) error {
	log.Printf("started processing: %s\n", job.VideoID)

	ctx := context.Background()

	video, err := h.Videos.FindByID(ctx, job.VideoID)
	if err != nil {
		return err
	}

	now := time.Now()
	if _, err := h.Videos.Update(ctx, job.VideoID, models.UpdateVideo{
		Status:      &models.STATUS_PROCESSING,
		ProcessedAt: &now,
	}); err != nil {
		return err
	}

	metadata, err := h.Probe.Analyze(video.OriginalPath)
	if err != nil {
		_, err := h.Videos.Update(ctx, job.VideoID, models.UpdateVideo{
			Status: &models.STATUS_FAILED,
		})
		if err != nil {
			return err
		}
	}

	if _, err := h.Videos.Update(ctx, job.VideoID, models.UpdateVideo{
		DurationSeconds: &metadata.Duration,
		Width:           &metadata.Width,
		Height:          &metadata.Height,
	}); err != nil {
		return err
	}

	return nil
}
