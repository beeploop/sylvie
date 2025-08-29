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

	jobMsgs, err := rabbit.ConnectToTranscodingQueue(cfg.TranscodingQueueName)
	if err != nil {
		log.Fatalf("Failed to connect to the transcoding job queue. Error: %s\n", err.Error())
	}

	if err := rabbit.ConnectToPublishQueue(cfg.PublishingQueueName); err != nil {
		log.Fatalf("Failed to connect to publishing job queue. Error: %s\n", err.Error())
	}

	go func() {
		for job := range jobMsgs {
			log.Printf("received a new msg\n")

			var msg rabbitmq.Message
			if err := json.Unmarshal(job.Body, &msg); err != nil {
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
				OutDir: filepath.Join(cfg.OutDir, "transcoded"),
			}

			metadata, err := transcoder.Transcode(params)
			if err != nil {
				log.Printf("Failed to transcode. Error: %s\n", err.Error())
				continue
			}

			if err := repo.Save(metadata); err != nil {
				log.Printf("Failed to save transcoding result. Error: %s\n", err.Error())
				continue
			}

			jsonData, err := json.Marshal(metadata)
			if err != nil {
				log.Printf("Failed to marshal the transcoding result. Error: %s\n", err.Error())
				continue
			}
			if err := rabbit.Publish(cfg.PublishingQueueName, jsonData); err != nil {
				log.Printf("Failed to publish the transcoding result. Error: %s\n", err.Error())
				continue
			}

			job.Ack(false)
		}
	}()

	fmt.Printf("Consuming messages... Press CTRL + C to stop\n")
	<-quitChan

	fmt.Printf("Shutting down gracefully...\n")
	close(quitChan)
	rabbit.Close()
}
