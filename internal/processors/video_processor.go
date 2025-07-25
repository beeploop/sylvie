package processors

import "github.com/beeploop/sylvie/internal/models"

type VideoProcessor struct{}

func (p *VideoProcessor) Process(job *models.Job, variant *models.Variant) (*ProcessResult, error) {
	return nil, nil
}
