package processors

import "github.com/beeploop/sylvie/internal/models"

type ImageProcessor struct {
	savePath string
}

func NewImageProcessor(savePath string) *ImageProcessor {
	return &ImageProcessor{
		savePath: savePath,
	}
}

func (p *ImageProcessor) Process(job *models.Job, variants []*models.Variant) ([]*ProcessResult, error) {
	return nil, nil
}
