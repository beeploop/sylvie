package transcoder

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/beeploop/sylvie/internal/metadata"
)

func createHSLMasterPlaylist(outFile string, rendetions []metadata.Rendetion) (metadata.Streaming, error) {
	f, err := os.Create(outFile)
	if err != nil {
		return metadata.Streaming{}, err
	}
	defer f.Close()

	if _, err := f.WriteString("#EXTM3U\n"); err != nil {
		return metadata.Streaming{}, err
	}
	if _, err := f.WriteString("#EXT-X-VERSION:3\n"); err != nil {
		return metadata.Streaming{}, err
	}

	for _, rendetion := range rendetions {
		resolution := fmt.Sprintf("%dx%d", rendetion.Width, rendetion.Height)
		variant := fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n", rendetion.Bitrate, resolution)
		if _, err := f.WriteString(variant); err != nil {
			return metadata.Streaming{}, err
		}

		baseDir := filepath.Dir(outFile)
		path := strings.TrimPrefix(rendetion.Outputs.HLS.Playlist, baseDir)
		cleanPath := strings.TrimPrefix(path, "/")
		if _, err := fmt.Fprintf(f, "%s\n\n", cleanPath); err != nil {
			return metadata.Streaming{}, err
		}
	}

	master := metadata.Streaming{
		MasterPlaylist: outFile,
		Protocol:       "hls",
	}

	return master, nil
}

func createHLSSegments(src, dest string, resolution Resolution) (metadata.HLSRendetionPath, error) {
	filename := fmt.Sprintf("%s.m3u8", resolution.Name())
	outFile := filepath.Join(dest, filename)

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i",
		src,
		"-profile:v",
		"baseline",
		"-level",
		"3.0",
		"-start_number",
		"0",
		"-hls_time",
		"10",
		"-hls_list_size",
		"0",
		"-f",
		"hls",
		outFile,
	)

	if err := cmd.Run(); err != nil {
		return metadata.HLSRendetionPath{}, err
	}

	segments, err := extractSegmentsFromM3u8(outFile)
	if err != nil {
		return metadata.HLSRendetionPath{}, err
	}

	hls := metadata.HLSRendetionPath{
		Playlist: outFile,
		Segments: segments,
	}

	return hls, nil
}

func extractSegmentsFromM3u8(m3u8Path string) ([]string, error) {
	f, err := os.Open(m3u8Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	baseDir := filepath.Dir(m3u8Path)

	var segments []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		segments = append(segments, filepath.Join(baseDir, line))
	}

	return segments, nil
}
