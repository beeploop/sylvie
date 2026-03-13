package application

import (
	"log"
	"sylvie/internal/config"
	"sylvie/internal/queue"
	"sylvie/internal/router/controllers"
	"sylvie/internal/storage"
)

func Bootstrap() *Application {
	conn, ch, err := queue.Connect(config.Load().Queue.ConnectionString)
	if err != nil {
		log.Fatalf("failed to open connection to rabbitmq: %s\n", err)
	}

	if err := queue.DeclareQueue(ch, config.Load().Queue.Name); err != nil {
		log.Fatalf("failed to declare transcoding queue: %s\n", err)
	}

	publisher := queue.NewPublisher(ch, config.Load().Queue.Name)

	store := storage.NewDiskStorage(storage.DiskStorageConfig{
		BaseDir:    config.Load().Storage.BaseDir,
		Permission: 0777,
	})

	uploadController := controllers.NewUploadControllerImpl(store)

	return &Application{
		RabbitConnection: conn,
		RabbitChannel:    ch,
		Publisher:        publisher,

		UploadController: uploadController,
	}
}
