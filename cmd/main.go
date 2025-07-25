package main

import (
	"log"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/beeploop/sylvie/internal/orchestrator"
	"github.com/beeploop/sylvie/internal/processors"
	"github.com/beeploop/sylvie/internal/source"
	"github.com/beeploop/sylvie/internal/storage"
)

func main() {
	path := "/home/screamour/repos/go/media-transcoding/config/config.yaml"

	cfg, err := config.Read(path)
	if err != nil {
		log.Fatalf("Error loading config: %s\n", err.Error())
	}

	jobSource := source.NewCLISource()
	imageProcessor := processors.ImageProcessor{}
	videoProcessor := processors.VideoProcessor{}
	diskStorage := storage.NewDiskStorage()

	maestro := orchestrator.NewOrchestrator(
		cfg,
		jobSource,
		imageProcessor,
		videoProcessor,
		diskStorage,
	)

	if err := maestro.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
