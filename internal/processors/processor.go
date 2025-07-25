package processors

import "github.com/beeploop/sylvie/internal/models"

type Processor interface {
	Process(job *models.Job, variants []*models.Variant) ([]*ProcessResult, error)
}
