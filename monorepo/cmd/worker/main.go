package main

import (
	"log"
	"os"
	"os/signal"
	"sylvie/internal/config"
	"sylvie/internal/queue"
	"syscall"
)

func main() {
	conn, ch, err := queue.Connect(config.Load().ConnectionString)
	if err != nil {
		log.Fatalf("failed to open connection to rabbitmq: %s\n", err)
	}

	if err := queue.DeclareQueue(ch, config.Load().QueueName); err != nil {
		log.Fatalf("failed to declare transcoding queue: %s\n", err)
	}

	consumer := queue.NewConsumer(ch, config.Load().QueueName)

	go func() {
		log.Println("starting rabbitmq consumer")
		if err := consumer.Consume(messageHandler); err != nil {
			log.Fatalf("failed to read message from rabbitmq: %s\n", err)
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT)
	<-quitChan

	log.Println("gracefully shutting down worker...")
	if err := ch.Close(); err != nil {
		log.Fatalf("error closing rabbitmq channel: %s\n", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("error closing rabbitmq connection: %s\n", err)
	}

	log.Println("worker exited")
}

func messageHandler(job queue.Job) error {
	log.Println("handler received a new job")
	log.Println(job)
	return nil
}
