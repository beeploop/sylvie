package controllers

import "sylvie/internal/video/entities"

type VideosController interface {
	Search(string) ([]entities.Video, error)
	Get(string) (entities.Video, error)
}
