package processors

import "github.com/beeploop/sylvie/internal/models"

type ProcessResult struct {
	Variant  models.Variant
	FilePath string
	Metadata models.MetaData
}
