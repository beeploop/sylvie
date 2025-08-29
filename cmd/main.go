package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"slices"
	"syscall"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/beeploop/sylvie/internal/rabbitmq"
	"github.com/beeploop/sylvie/internal/repository"
	"github.com/beeploop/sylvie/internal/transcoder"
	"github.com/beeploop/sylvie/internal/utils"
)

func main() {
	configFile := flag.String("config", "", "Specify the yaml configuration file")
	flag.Parse()

	cfg := config.Init(configFile)

	repo := repository.NewDiskRepository(filepath.Join(cfg.OutDir, "results"))
	rabbit := rabbitmq.Init(cfg)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for d := range rabbit.Msgs {
			log.Printf("received a new msg\n")

			var msg rabbitmq.Message
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Error reading msg body: %s\n", err.Error())
				continue
			}

			params := &transcoder.TranscodeInput{
				VideoID: msg.VideoID,
				InFile:  msg.Path,
				Resolutions: slices.Collect(utils.Map(
					msg.Resolutions,
					func(resolution string) transcoder.Resolution {
						return transcoder.ResolutionFromName(resolution)
					},
				)),
				OutDir: filepath.Join(cfg.OutDir, "encoded"),
			}

			metadata, err := transcoder.Transcode(params)
			if err != nil {
				log.Printf("Failed to transcode. Error: %s\n", err.Error())
				continue
			}

			if err := repo.Save(metadata); err != nil {
				log.Printf("Failed to save encoding result. Error: %s\n", err.Error())
				continue
			}

			d.Ack(false)
		}
	}()

	fmt.Printf("Consuming messages... Press CTRL + C to stop\n")
	<-quitChan

	fmt.Printf("Shutting down gracefully...\n")
	rabbit.Close()
	close(quitChan)
}
