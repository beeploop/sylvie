package controllers

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"
	"sylvie/internal/storage"
	"sylvie/internal/video/entities"
	"sylvie/internal/video/models"
	"sylvie/internal/video/repositories"

	"github.com/google/uuid"
)

type uploadControllerImpl struct {
	videos repositories.VideoRepository
	store  storage.Storage
}

func NewUploadControllerImpl(videos repositories.VideoRepository, fs storage.Storage) *uploadControllerImpl {
	return &uploadControllerImpl{
		videos: videos,
		store:  fs,
	}
}

func (c *uploadControllerImpl) Upload(file *multipart.FileHeader, title string) (entities.NewVideo, error) {
	var video entities.NewVideo

	if file == nil || title == "" {
		return video, errors.New("missing required data")
	}

	content, err := file.Open()
	if err != nil {
		return video, err
	}
	defer content.Close()

	data, err := io.ReadAll(content)
	if err != nil {
		return video, err
	}

	videoID := uuid.NewString()
	subpath := filepath.Join(videoID, "original.mp4")

	fullpath, err := c.store.Write(context.Background(), subpath, data)
	if err != nil {
		return video, err
	}

	videoData := models.NewVideo{
		ID:           videoID,
		Title:        title,
		OriginalPath: fullpath,
		Status:       models.STATUS_UPLOADED,
	}
	return c.videos.Create(context.Background(), videoData)
}
