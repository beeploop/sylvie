package application

import (
	"sylvie/internal/http/controllers"
	"sylvie/internal/queue"
	"sylvie/internal/transcoding"
	"sylvie/internal/video/repositories"

	"github.com/streadway/amqp"
)

type APIApplication struct {
	RabbitConnection *amqp.Connection
	RabbitChannel    *amqp.Channel
	Publisher        queue.Publisher

	UploadController controllers.UploadController
	VideosController controllers.VideosController
}

type WorkerApplication struct {
	Videos     repositories.VideoRepository
	Probe      transcoding.Probe
	Transcoder transcoding.Transcoder
	Thumbnails transcoding.ThumbnailGenerator
	Playlist   transcoding.PlaylistGenerator
}
