package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/beeploop/sylvie/internal/transcoder"
)

func main() {
	inputFile := "/home/screamour/Videos/unwrapped-beeploop.mp4"
	outDir := "/home/screamour/repos/go/media-transcoding"

	resolutions := []transcoder.Resolution{
		transcoder.RES_1080p,
		transcoder.RES_720p,
		transcoder.RES_480p,
		transcoder.RES_360p,
		transcoder.RES_240p,
		transcoder.RES_144p,
	}

	doneChan := make(chan string)
	var wg sync.WaitGroup

	for _, resolution := range resolutions {
		wg.Add(1)

		go func() {
			defer wg.Done()

			fmt.Printf("transcoding input to resolution: %s\n", resolution.Name())

			output := filepath.Join(outDir, fmt.Sprintf("%s.mp4", resolution.Name()))
			err := transcoder.Transcode(inputFile, output, transcoder.TemplateFactory(resolution))
			if err != nil {
				log.Fatal(err.Error())
			}

			doneChan <- fmt.Sprintf("transcoded input to resolution: %s", resolution.Name())
		}()
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	for msg := range doneChan {
		fmt.Println(msg)
	}

	fmt.Println("done")
}
