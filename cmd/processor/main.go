package main

import (
	"log"
	"os"
	"os/signal"
	"sylvie/internal/config"
	"sylvie/internal/queue"
	"sylvie/internal/workers"
	"syscall"
)

func main() {
	conn, ch, err := queue.Connect(config.Load().Queue.ConnectionString)
	if err != nil {
		log.Fatalf("failed to open connection to rabbitmq: %s\n", err)
	}

	if err := queue.DeclareQueue(ch, config.Load().Queue.Name); err != nil {
		log.Fatalf("failed to declare transcoding queue: %s\n", err)
	}

	consumer := queue.NewConsumer(ch, config.Load().Queue.Name)
	manager := workers.NewManager(config.Load())

	errChan := make(chan error, 1)
	go func() {
		log.Println("starting rabbitmq consumer")
		if err := consumer.Consume(manager.Handle); err != nil {
			errChan <- err
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-quitChan:
		log.Printf("received a shutdown signal: %s\n", sig)

	case err := <-errChan:
		log.Printf("worker encountered an error: %s\n", err)
	}

	log.Println("gracefully shutting down worker...")

	ch.Close()
	conn.Close()

	log.Println("worker exited")
}
