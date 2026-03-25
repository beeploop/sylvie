package controllers

import (
	"context"
	"sylvie/internal/video/entities"
	"sylvie/internal/video/repositories"
)

type videosControllerImpl struct {
	videos repositories.VideoRepository
}

func NewVideosControllerImpl(videos repositories.VideoRepository) *videosControllerImpl {
	return &videosControllerImpl{
		videos: videos,
	}
}

func (c *videosControllerImpl) Search(title string) ([]entities.Video, error) {
	return c.videos.FindByTitle(context.Background(), title)
}

func (c *videosControllerImpl) Get(id string) (entities.Video, error) {
	return c.videos.FindByID(context.Background(), id)
}
