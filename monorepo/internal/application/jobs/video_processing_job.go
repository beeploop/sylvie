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
	ctx := context.Background()

	_, err := h.Videos.FindByID(ctx, job.VideoID)
	if err != nil {
		return err
	}

	now := time.Now()
	h.Videos.Update(ctx, job.VideoID, models.UpdateVideo{
		Status:      &models.STATUS_PROCESSING,
		ProcessedAt: &now,
	})

	log.Printf("started processing: %s\n", job.VideoID)
	return nil
}
