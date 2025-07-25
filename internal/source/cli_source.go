package source

import (
	"errors"
	"time"

	"github.com/beeploop/sylvie/internal/models"
	"github.com/google/uuid"
)

type CLISource struct {
	sources []string
	index   int
}

func NewCLISource() *CLISource {
	return &CLISource{
		sources: make([]string, 0),
		index:   0,
	}
}

func (s *CLISource) NextJob() (*models.Job, error) {
	if !s.HasNext() {
		return nil, errors.New("no more jobs")
	}

	source := s.sources[s.index]
	s.index++

	job := &models.Job{
		ID:        uuid.New().String(),
		MediaType: detectMediaType(source),
		FilePath:  source,
		CreatedAt: time.Now(),
	}

	return job, nil
}

func (s *CLISource) HasNext() bool {
	return s.index < len(s.sources)
}

func (s *CLISource) Queue(source string) {
	s.sources = append(s.sources, source)
}
