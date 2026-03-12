package transcoding

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type transcoderImpl struct {
	BaseDir    string
	Permission os.FileMode
}

func NewTranscoder(baseDir string, permission os.FileMode) *transcoderImpl {
	return &transcoderImpl{
		BaseDir:    baseDir,
		Permission: permission,
	}
}

func (t *transcoderImpl) Transcode(rendetion Rendetion) (string, error) {
	outDir := t.outputDirectory(rendetion.VideoID, rendetion.Resolution.Name())
	if err := os.MkdirAll(outDir, t.Permission); err != nil {
		return "", err
	}

	cmd := t.buildCommand(rendetion, outDir)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/index.m3u8", outDir), nil
}

func (t *transcoderImpl) outputDirectory(videoID, resolution string) string {
	return filepath.Join(t.BaseDir, videoID, resolution)
}

func (t *transcoderImpl) buildCommand(rendetion Rendetion, outDir string) *exec.Cmd {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", rendetion.InputPath,
		"-vf", "scale="+rendetion.Resolution.Ratio(),
		"-c:v", "libx264", "-b:v", ToBitrateSuffixNotation(rendetion.Resolution.VideoBitRate()),
		"-c:a", "aac", "-b:a", ToBitrateSuffixNotation(rendetion.Resolution.AudioBitRate()),
		"-f", "hls",
		"-hls_time", "5",
		"-hls_playlist_type", "vod",
		"-hls_segment_filename", fmt.Sprintf("%s/segment_%%03d.ts", outDir),
		fmt.Sprintf("%s/index.m3u8", outDir),
	)

	return cmd
}
