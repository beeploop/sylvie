package controllers

import (
	"mime/multipart"
	"sylvie/internal/http/dtos"
)

type UploadController interface {
	Upload(*multipart.FileHeader, string) (dtos.UploadResult, error)
}
