package application

import (
	"os"
	"sylvie/internal/config"
	"sylvie/internal/transcoding"
	"sylvie/internal/video/repositories"
)

func BootstrapWorker() *WorkerApplication {
	videoRepository := repositories.NewVideoJSONFileRepository("tmp/db.json", os.FileMode(0777))

	probe := transcoding.NewProbeImpl(config.Load().FFMPEG.FfprobePath)
	transcoder := transcoding.NewTranscoder(config.Load().FFMPEG.FfmpegPath, "tmp/transcoded/", os.FileMode(0777))
	thumbnailGenerator := transcoding.NewThumbnailGenerator(config.Load().FFMPEG.FfmpegPath, "tmp/transcoded/", os.FileMode(0777))
	playlistGenerator := transcoding.NewPlaylistGenerator("tmp/transcoded/", os.FileMode(0777))

	return &WorkerApplication{
		Videos:     videoRepository,
		Probe:      probe,
		Transcoder: transcoder,
		Thumbnails: thumbnailGenerator,
		Playlist:   playlistGenerator,
	}
}
