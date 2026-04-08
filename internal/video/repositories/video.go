package repositories

import (
	"context"
	"sylvie/internal/video/entities"
	"sylvie/internal/video/models"
)

type VideoRepository interface {
	Create(context.Context, models.NewVideo) (entities.NewVideo, error)
	FindByTitle(context.Context, string) ([]entities.Video, error)
	FindByID(context.Context, string) (entities.Video, error)
	Update(context.Context, string, models.UpdateVideo) (entities.Video, error)
}
