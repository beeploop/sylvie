package transcoder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/beeploop/sylvie/internal/metadata"
	"github.com/beeploop/sylvie/internal/utils"
	"github.com/google/uuid"
)

type TranscodeInput struct {
	VideoID     string
	InFile      string
	Resolutions []Resolution
	Config      *config.Config
}

func Transcode(params *TranscodeInput) (metadata.Metadata, error) {
	var wg sync.WaitGroup

	meta := metadata.Metadata{
		VideoID:    params.VideoID,
		Rendetions: make([]metadata.Rendetion, 0),
		Thumbnails: make([]metadata.Thumbnail, 0),
	}

	dest := filepath.Join(params.Config.OutDir, uuid.NewString())
	if err := os.MkdirAll(dest, 0777); err != nil {
		return meta, err
	}

	result, err := extractDataWithFfprobe(params.InFile)
	if err != nil {
		return meta, err
	}

	sourceMetadata, err := result.toSourceMetadata()
	if err != nil {
		return meta, err
	}
	meta.SourceMetadata = sourceMetadata

	for _, resolution := range params.Resolutions {
		wg.Add(1)

		go func() {
			defer wg.Done()

			template := TemplateFactory(resolution)
			output := filepath.Join(dest, fmt.Sprintf("%s.mp4", resolution.Name()))

			rendetion, err := createRendetion(params.InFile, output, sourceMetadata, template)
			if err != nil {
				log.Printf("Failed creating %s rendetion. Error: %s\n", resolution.Name(), err.Error())
				return
			}
			meta.Rendetions = append(meta.Rendetions, rendetion)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		seektime := 5
		output := filepath.Join(dest, "thumbnail_default.jpg")
		thumbnail, err := createThumbnail(params.InFile, output, seektime)
		if err != nil {
			log.Printf("Failed to create thumbnail: %s\n", err.Error())
			return
		}
		meta.Thumbnails = append(meta.Thumbnails, thumbnail)
	}()

	wg.Wait()
	return meta, nil
}

func createRendetion(src, dest string, meta metadata.SourceMetadata, template Template) (metadata.Rendetion, error) {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i",
		src,
		"-vf",
		fmt.Sprintf("scale=%d:%d", template.Width, template.Height),
		"-preset",
		template.Preset,
		"-crf",
		strconv.Itoa(template.CRF),
		"-c:v",
		template.VideoCodec,
		"-b:v",
		utils.BitrateToFfmpegStyleString(template.VideoBitRate),
		"-c:a",
		template.AudioCodec,
		"-b:a",
		utils.BitrateToFfmpegStyleString(template.AudioBitRate),
		dest,
	)

	if err := cmd.Run(); err != nil {
		return metadata.Rendetion{}, err
	}

	info, err := os.Stat(dest)
	if err != nil {
		return metadata.Rendetion{}, err
	}

	rendetion := metadata.Rendetion{
		Resolution: template.Resolution.Name(),
		Width:      template.Width,
		Height:     template.Height,
		FrameRate:  meta.FrameRate,
		Bitrate:    template.VideoBitRate,
		Codec:      template.VideoCodec,
		FileSize:   int(info.Size()),
		Path:       dest,
	}
	return rendetion, nil
}

func createThumbnail(src, dest string, seektimeInSeconds int) (metadata.Thumbnail, error) {
	timestampString := utils.SecondsToTimestamp(seektimeInSeconds)
	timestampInt, err := utils.TimestampToSeconds(timestampString)
	if err != nil {
		return metadata.Thumbnail{}, err
	}

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i",
		src,
		"-ss",
		timestampString,
		"-vframes",
		"1",
		"-q:v",
		"2",
		dest,
	)

	if err := cmd.Run(); err != nil {
		return metadata.Thumbnail{}, err
	}

	return metadata.Thumbnail{
		Timestamp: timestampInt,
		Type:      "default",
		Path:      dest,
	}, nil
}
