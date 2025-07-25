package processors

import "github.com/beeploop/sylvie/internal/models"

type Processor interface {
	Process(job *models.Job, variant *models.Variant) (*ProcessResult, error)
}
