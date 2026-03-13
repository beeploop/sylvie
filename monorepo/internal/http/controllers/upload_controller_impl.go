package controllers

import (
	"context"
	"io"
	"mime/multipart"
	"path/filepath"
	"sylvie/internal/http/dtos"
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

func (c *uploadControllerImpl) Upload(file *multipart.FileHeader, title string) (dtos.UploadResult, error) {
	var result dtos.UploadResult

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
	subpath := filepath.Join(videoID, "original.mp4")

	fullpath, err := c.store.Write(context.Background(), subpath, data)
	if err != nil {
		return result, err
	}

	result.VideoID = videoID
	result.Status = "processing"
	result.Path = fullpath

	return result, nil
}
