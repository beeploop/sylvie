package main

import (
	"fmt"
	"log"
	"path/filepath"
	"slices"
	"sync"

	"github.com/beeploop/sylvie/internal/transcoder"
	"github.com/beeploop/sylvie/internal/utils"
)

func main() {
	inputFile := "/home/screamour/Videos/unwrapped-beeploop.mp4"
	outDir := "/home/screamour/repos/go/media-transcoding"

	resolutions := slices.Collect(utils.Map(
		[]string{"1080p", "720p", "480p", "360p", "240p", "144p"},
		func(res string) transcoder.Resolution {
			return transcoder.ResolutionFromName(res)
		},
	))

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
