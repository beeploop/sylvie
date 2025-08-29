package main

import (
	"flag"
	"log"
	"slices"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/beeploop/sylvie/internal/repository"
	"github.com/beeploop/sylvie/internal/transcoder"
	"github.com/beeploop/sylvie/internal/utils"
	"github.com/google/uuid"
)

func main() {
	configFile := flag.String("config", "", "Specify the yaml configuration file")
	flag.Parse()

	cfg := config.Init(configFile)

	inputFile := "/home/screamour/Videos/unwrapped-beeploop.mp4"
	resolutions := slices.Collect(utils.Map(
		[]string{"1080p", "720p", "480p", "360p", "240p", "144p"},
		func(res string) transcoder.Resolution {
			return transcoder.ResolutionFromName(res)
		},
	))

	params := &transcoder.TranscodeInput{
		VideoID:     uuid.NewString(),
		InFile:      inputFile,
		Resolutions: resolutions,
		Config:      cfg,
	}

	metadata, err := transcoder.Transcode(params)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := repository.NewDiskRepository(cfg.OutDir)
	if err := repo.Save(metadata); err != nil {
		log.Fatal(err.Error())
	}
}
