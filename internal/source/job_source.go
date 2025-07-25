package source

import "github.com/beeploop/sylvie/internal/models"

type JobSource interface {
	NextJob() (*models.Job, error)
	HasNext() bool
}
