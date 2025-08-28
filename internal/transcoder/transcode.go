package transcoder

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/beeploop/sylvie/internal/metadata"
	"github.com/beeploop/sylvie/internal/utils"
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

	dest := filepath.Join(params.Config.OutDir, meta.VideoID)
	if err := os.MkdirAll(dest, 0777); err != nil {
		return meta, err
	}

	result, err := extractMetadataWithFfprobe(params.InFile)
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

			outDir := filepath.Join(dest, resolution.Name())
			if err := os.MkdirAll(outDir, 0777); err != nil {
				log.Printf("Failed to create a directory for %s, rendetion: Error: %s\n", resolution.Name(), err.Error())
				return
			}

			template := TemplateFactory(resolution)
			rendetionOutFile := filepath.Join(outDir, fmt.Sprintf("%s.mp4", resolution.Name()))

			rendetion, err := createRendetion(
				params.InFile,
				rendetionOutFile,
				sourceMetadata,
				template)
			if err != nil {
				log.Printf("Failed creating %s rendetion. Error: %s\n", resolution.Name(), err.Error())
				return
			}

			hlsOutDir := filepath.Dir(rendetion.Outputs.MP4.Path)
			hlsRendetion, err := createHLS(rendetion.Outputs.MP4.Path, hlsOutDir, resolution)
			if err != nil {
				log.Printf("Failed creating HLS for %s rendetion. Error: %s\n", resolution.Name(), err.Error())
			}

			rendetion.Outputs.HLS = hlsRendetion
			meta.Rendetions = append(meta.Rendetions, rendetion)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		seektime := 5
		thumbnailOutFile := filepath.Join(dest, "thumbnail_default.jpg")
		thumbnail, err := createThumbnail(params.InFile, thumbnailOutFile, seektime)
		if err != nil {
			log.Printf("Failed to create thumbnail: %s\n", err.Error())
			return
		}
		meta.Thumbnails = append(meta.Thumbnails, thumbnail)
	}()

	wg.Wait()

	masterPlaylistOutFile := filepath.Join(dest, "master.m3u8")
	masterPlaylist, err := createHLSMasterPlaylist(masterPlaylistOutFile, meta.Rendetions)
	if err != nil {
		return meta, err
	}
	meta.Streaming = masterPlaylist

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
		Outputs: metadata.RendetionOutputs{
			MP4: metadata.MP4RendetionPath{
				Path: dest,
			},
		},
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

func createHLSMasterPlaylist(outFile string, rendetions []metadata.Rendetion) (metadata.Streaming, error) {
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
		if _, err := f.WriteString(fmt.Sprintf("%s\n\n", cleanPath)); err != nil {
			return metadata.Streaming{}, err
		}
	}

	master := metadata.Streaming{
		MasterPlaylist: outFile,
		Protocol:       "hls",
	}

	return master, nil
}

func createHLS(src, dest string, resolution Resolution) (metadata.HLSRendetionPath, error) {
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

	segments, err := extractSegments(outFile)
	if err != nil {
		return metadata.HLSRendetionPath{}, err
	}

	hls := metadata.HLSRendetionPath{
		Playlist: outFile,
		Segments: segments,
	}

	return hls, nil
}

func extractSegments(m3u8Path string) ([]string, error) {
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
