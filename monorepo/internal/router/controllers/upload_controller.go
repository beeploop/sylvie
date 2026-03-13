package controllers

import "mime/multipart"

type UploadResultDTO struct {
	VideoID string `json:"video_id"`
	Status  string `json:"status"`
}

type UploadController interface {
	Upload(*multipart.FileHeader, string) (UploadResultDTO, error)
}
