package processors

import "github.com/beeploop/sylvie/internal/models"

type ImageProcessor struct{}

func (p *ImageProcessor) Process(job *models.Job, variant *models.Variant) (*ProcessResult, error) {
	return nil, nil
}
