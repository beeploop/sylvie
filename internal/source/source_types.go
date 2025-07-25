package source

import (
	"path/filepath"
	"slices"

	"github.com/beeploop/sylvie/internal/models"
)

func detectMediaType(source string) models.MediaType {
	imageTypes := []string{".jpg", ".jpeg", ".png", ".webp"}
	videoTypes := []string{".mp4", ".mov", ".wmv", ".webm"}

	ext := filepath.Ext(source)

	if slices.Contains(imageTypes, ext) {
		return models.IMAGE
	}

	if slices.Contains(videoTypes, ext) {
		return models.VIDEO
	}

	return models.UNKNOWN
}
