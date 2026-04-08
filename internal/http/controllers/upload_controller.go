package controllers

import (
	"mime/multipart"
	"sylvie/internal/video/entities"
)

type UploadController interface {
	Upload(*multipart.FileHeader, string) (entities.NewVideo, error)
}
