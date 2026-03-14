package controllers

import "sylvie/internal/video/entities"

type VideosController interface {
	Get(string) (entities.Video, error)
}
