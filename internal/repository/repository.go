package repository

import "github.com/beeploop/sylvie/internal/metadata"

type Repository interface {
	Save(metadata.Metadata) error
}
