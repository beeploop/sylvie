package processors

import "github.com/beeploop/sylvie/internal/models"

type ProcessResult struct {
	Variant    models.Variant
	OutputPath string
	Metadata   models.MetaData
}
