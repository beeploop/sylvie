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

func BootstrapAPI() *APIApplication {
	conn, ch, err := queue.Connect(config.Load().Queue.ConnectionString)
	if err != nil {
		log.Fatalf("failed to open connection to rabbitmq: %s\n", err)
	}

	if err := queue.DeclareQueue(ch, config.Load().Queue.Name); err != nil {
		log.Fatalf("failed to declare transcoding queue: %s\n", err)
	}

	publisher := queue.NewPublisher(ch, config.Load().Queue.Name)

	videoRepository := repositories.NewVideoJSONFileRepository("tmp/db.json", os.FileMode(0777))

	store := storage.NewDiskStorage(storage.DiskStorageConfig{
		BaseDir:    config.Load().Storage.BaseDir,
		Permission: 0777,
	})

	uploadController := controllers.NewUploadControllerImpl(videoRepository, store)
	videosController := controllers.NewVideosControllerImpl(videoRepository)

	return &APIApplication{
		RabbitConnection: conn,
		RabbitChannel:    ch,
		Publisher:        publisher,

		UploadController: uploadController,
		VideosController: videosController,
	}
}
