package orchestrator

import (
	"fmt"
	"log"
	"slices"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/beeploop/sylvie/internal/models"
	"github.com/beeploop/sylvie/internal/pkg/utils"
	"github.com/beeploop/sylvie/internal/processors"
	"github.com/beeploop/sylvie/internal/source"
	"github.com/beeploop/sylvie/internal/storage"
)

type Orchestrator struct {
	config         *config.Config
	jobSource      source.JobSource
	imageProcessor *processors.ImageProcessor
	videoProcessor *processors.VideoProcessor
	storage        storage.Storage
}

func NewOrchestrator(
	config *config.Config,
	jobSource source.JobSource,
	imageProcessor *processors.ImageProcessor,
	videoProcessor *processors.VideoProcessor,
	storage storage.Storage,
) *Orchestrator {
	return &Orchestrator{
		config:         config,
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

		results := make([]*processors.ProcessResult, 0)
		switch job.MediaType {
		case models.IMAGE:
			variants := slices.AppendSeq(
				make([]*models.Variant, 0),
				utils.Map(
					o.config.Processing.Image.Variants,
					func(variant config.ImageVariants) *models.Variant {
						return &models.Variant{
							Name:   variant.Name,
							Width:  variant.Width,
							Height: variant.Height,
						}
					},
				),
			)

			transcodings, err := o.imageProcessor.Process(job, variants)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			results = transcodings

		case models.VIDEO:
			variants := slices.AppendSeq(
				make([]*models.Variant, 0),
				utils.Map(
					o.config.Processing.Video.Variants,
					func(variant config.VideoVariants) *models.Variant {
						return &models.Variant{
							Name:         variant.Name,
							Width:        variant.Width,
							Height:       variant.Height,
							VideoBitRate: variant.VideoBitRate,
							AudioBitRate: variant.AudioBitRate,
						}
					},
				),
			)

			transcodings, err := o.videoProcessor.Process(job, variants)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			results = transcodings

		case models.UNKNOWN:
			log.Println("unknown media type")
			continue
		}

		fmt.Println(results)
	}

	return nil
}
