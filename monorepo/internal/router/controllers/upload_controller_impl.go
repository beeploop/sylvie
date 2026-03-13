package controllers

import (
	"context"
	"io"
	"mime/multipart"
	"path/filepath"
	"sylvie/internal/storage"

	"github.com/google/uuid"
)

type uploadControllerImpl struct {
	store storage.Storage
}

func NewUploadControllerImpl(fs storage.Storage) *uploadControllerImpl {
	return &uploadControllerImpl{
		store: fs,
	}
}

func (c *uploadControllerImpl) Upload(file *multipart.FileHeader, title string) (UploadResultDTO, error) {
	var result UploadResultDTO

	content, err := file.Open()
	if err != nil {
		return result, err
	}
	defer content.Close()

	data, err := io.ReadAll(content)
	if err != nil {
		return result, err
	}

	videoID := uuid.NewString()
	path := filepath.Join(videoID, "original.mp4")

	if _, err := c.store.Write(context.Background(), path, data); err != nil {
		return result, err
	}

	result.VideoID = videoID
	result.Status = "processing"

	return result, nil
}
