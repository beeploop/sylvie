package application

import (
	"log"
	"sylvie/internal/config"
	"sylvie/internal/queue"
)

func Bootstrap() *Application {
	conn, ch, err := queue.Connect(config.Load().ConnectionString)
	if err != nil {
		log.Fatalf("failed to open connection to rabbitmq: %s\n", err)
	}

	if err := queue.DeclareQueue(ch, config.Load().QueueName); err != nil {
		log.Fatalf("failed to declare transcoding queue: %s\n", err)
	}

	publisher := queue.NewPublisher(ch, config.Load().QueueName)

	return &Application{
		RabbitConnection: conn,
		RabbitChannel:    ch,
		Publisher:        publisher,
	}
}
