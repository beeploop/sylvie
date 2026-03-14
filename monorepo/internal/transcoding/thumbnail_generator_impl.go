package transcoding

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type thumbnailGeneratorImpl struct {
	FfmpegPath string
	BaseDir    string
	Permission os.FileMode
}

func NewThumbnailGenerator(ffmpegPath, baseDir string, permission os.FileMode) *thumbnailGeneratorImpl {
	return &thumbnailGeneratorImpl{
		FfmpegPath: ffmpegPath,
		BaseDir:    baseDir,
		Permission: permission,
	}
}

func (g *thumbnailGeneratorImpl) Generate(input ThumbnailInput) (string, error) {
	outDir := g.outputDirectory(input.VideoID)
	if err := os.MkdirAll(outDir, g.Permission); err != nil {
		return "", err
	}

	cmd := g.buildCommand(input.Filepath, outDir)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/thumbnail.jpg", outDir), nil
}

func (g *thumbnailGeneratorImpl) outputDirectory(videoID string) string {
	return filepath.Join(g.BaseDir, videoID)
}

func (g *thumbnailGeneratorImpl) buildCommand(input, outDir string) *exec.Cmd {
	cmd := exec.Command(
		g.FfmpegPath,
		"-ss", "3",
		"-i", input,
		"-frames:v", "1",
		fmt.Sprintf("%s/thumbnail.jpg", outDir),
	)

	return cmd
}
