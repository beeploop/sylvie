package orchestrator

import (
	"fmt"
	"log"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/beeploop/sylvie/internal/processors"
	"github.com/beeploop/sylvie/internal/source"
	"github.com/beeploop/sylvie/internal/storage"
)

type Orchestrator struct {
	config         *config.Config
	jobSource      source.JobSource
	imageProcessor processors.ImageProcessor
	videoProcessor processors.VideoProcessor
	storage        storage.Storage
}

func NewOrchestrator(
	config *config.Config,
	jobSource source.JobSource,
	imageProcessor processors.ImageProcessor,
	videoProcessor processors.VideoProcessor,
	storage storage.Storage,
) *Orchestrator {
	return &Orchestrator{
		jobSource:      jobSource,
		imageProcessor: imageProcessor,
		videoProcessor: videoProcessor,
		storage:        storage,
	}
}

func (o *Orchestrator) Run() error {
	for o.jobSource.HasNext() {
		job, err := o.jobSource.NextJob()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		fmt.Println(job)
	}

	return nil
}
