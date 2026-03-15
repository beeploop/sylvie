package entities

import "sylvie/internal/video/models"

type NewVideo struct {
	ID           string
	Title        string
	Status       string
	OriginalPath string
}

func ModelToNewVideo(model models.NewVideo) NewVideo {
	return NewVideo{
		ID:           model.ID,
		Title:        model.Title,
		Status:       string(model.Status),
		OriginalPath: model.OriginalPath,
	}
}

type Video struct {
	ID                 string
	Title              string
	Status             string
	OriginalPath       string
	MasterPlaylistPath string
	ThumbnailPath      string
	DurationSeconds    float64
	Width              int
	Height             int
}

func ModelToVideo(model models.Video) Video {
	return Video{
		ID:                 model.ID,
		Title:              model.Title,
		Status:             string(model.Status),
		OriginalPath:       model.OriginalPath,
		MasterPlaylistPath: model.MasterPlaylistPath.String,
		ThumbnailPath:      model.ThumbnailPath.String,
		DurationSeconds:    model.DurationSeconds.Float64,
		Width:              int(model.Width.Int64),
		Height:             int(model.Height.Int64),
	}
}
