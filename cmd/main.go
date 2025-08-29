package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
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

	rabbit := rabbitmq.Init(cfg)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for d := range rabbit.Msgs {
			log.Printf("received a new msg\n")

			var msg rabbitmq.Message
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Error reading msg body: %s\n", err.Error())
				return
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
				Config: cfg,
			}

			metadata, err := transcoder.Transcode(params)
			if err != nil {
				log.Fatal(err.Error())
			}

			b, err := json.MarshalIndent(metadata, "", "  ")
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(string(b))

			d.Ack(false)
		}
	}()

	fmt.Printf("Consuming messages... Press CTRL + C to stop\n")
	<-quitChan

	fmt.Printf("Shutting down gracefully...")
	rabbit.Close()
}
