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
	input := "/home/screamour/Videos/huhu.mp4"
	path := "/home/screamour/repos/go/media-transcoding/config/config.yaml"

	cfg, err := config.Read(path)
	if err != nil {
		log.Fatalf("Error loading config: %s\n", err.Error())
	}

	jobSource := source.NewCLISource()
	jobSource.Queue(input)

	imageProcessor := processors.NewImageProcessor(cfg.App.TempDir)
	videoProcessor := processors.NewVideoProcessor(
		cfg.App.TempDir,
		cfg.App.FfmpegPath,
		cfg.App.FfprobePath,
	)
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
