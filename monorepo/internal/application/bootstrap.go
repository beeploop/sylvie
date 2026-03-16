package application

import (
	"log"
	"os"
	"sylvie/internal/config"
	"sylvie/internal/http/controllers"
	"sylvie/internal/queue"
	"sylvie/internal/storage"
	"sylvie/internal/video/repositories"
)

func Bootstrap(config *config.Config) *Application {
	conn, ch, err := queue.Connect(config.Queue.ConnectionString)
	if err != nil {
		log.Fatalf("failed to open connection to rabbitmq: %s\n", err)
	}

	if err := queue.DeclareQueue(ch, config.Queue.Name); err != nil {
		log.Fatalf("failed to declare transcoding queue: %s\n", err)
	}

	publisher := queue.NewPublisher(ch, config.Queue.Name)

	videoRepository := repositories.NewVideoJSONFileRepository("tmp/db.json", os.FileMode(0777))

	store := storage.NewDiskStorage(storage.DiskStorageConfig{
		BaseDir:    config.Storage.BaseDir,
		Permission: 0777,
	})

	uploadController := controllers.NewUploadControllerImpl(videoRepository, store)
	videosController := controllers.NewVideosControllerImpl(videoRepository)

	return &Application{
		RabbitConnection: conn,
		RabbitChannel:    ch,
		Publisher:        publisher,

		UploadController: uploadController,
		VideosController: videosController,
	}
}
