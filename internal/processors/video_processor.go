package processors

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/beeploop/sylvie/internal/models"
)

type VideoProcessor struct {
	savePath    string
	ffmpegPath  string
	ffprobePath string
}

func NewVideoProcessor(savePath, ffmpegPath, ffprobePath string) *VideoProcessor {
	return &VideoProcessor{
		savePath:    savePath,
		ffmpegPath:  ffmpegPath,
		ffprobePath: ffprobePath,
	}
}

func (p *VideoProcessor) Process(job *models.Job, variants []*models.Variant) ([]*ProcessResult, error) {
	outputDirectory := filepath.Join(p.savePath, job.ID)
	extension := filepath.Ext(job.FilePath)

	if err := os.MkdirAll(outputDirectory, 0777); err != nil {
		return nil, err
	}

	for _, variant := range variants {
		outputFilename := fmt.Sprintf("output_%s%s", variant.Name, extension)

		cmd := exec.Command(
			"ffmpeg",
			"-i",
			job.FilePath,
			"-vf",
			fmt.Sprintf("scale=%d:%d", variant.Width, variant.Height),
			"-c:v",
			"libx264",
			"-b:v",
			variant.VideoBitRate,
			"-c:a",
			"libmp3lame",
			"-b:a",
			variant.AudioBitRate,
			filepath.Join(outputDirectory, outputFilename),
		)

		if err := cmd.Run(); err != nil {
			return nil, err
		}
	}

	if err := p.extractAudio(job, variants[0], outputDirectory); err != nil {
		return nil, err
	}

	return nil, nil
}

func (p *VideoProcessor) extractAudio(job *models.Job, variant *models.Variant, outDir string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		job.FilePath,
		"-vn",
		"-b:a",
		variant.AudioBitRate,
		"-c:a",
		"libmp3lame",
		filepath.Join(outDir, "audio.mp3"),
	)

	return cmd.Run()
}
