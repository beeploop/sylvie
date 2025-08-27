package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/beeploop/sylvie/internal/transcoder"
	"github.com/beeploop/sylvie/internal/utils"
	"github.com/google/uuid"
)

func main() {
	inputFile := "/home/screamour/Videos/unwrapped-beeploop.mp4"
	outDir := "/home/screamour/repos/go/media-transcoding/results"

	if err := os.MkdirAll(outDir, 0777); err != nil {
		log.Fatal(err.Error())
	}

	resolutions := slices.Collect(utils.Map(
		[]string{"1080p", "720p", "480p", "360p", "240p", "144p"},
		func(res string) transcoder.Resolution {
			return transcoder.ResolutionFromName(res)
		},
	))

	params := &transcoder.TranscodeInput{
		VideoID:     uuid.NewString(),
		InFile:      inputFile,
		OutDir:      outDir,
		Resolutions: resolutions,
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
}
