package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sylvie/internal/application"
	"sylvie/internal/http"
	"syscall"
)

func main() {
	app := application.Bootstrap()

	server := http.NewServer(app)

	errChan := make(chan error, 1)
	go func() {
		log.Printf("api server listening in port: %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-quitChan:
		log.Printf("received a shutdown signal: %s\n", sig)

	case err := <-errChan:
		log.Printf("server encountered an error: %s\n", err)
	}

	log.Println("gracefully shutting down rabbitmq...")

	if err := app.RabbitChannel.Close(); err != nil {
		log.Fatalf("faild to close rabbitmq channel: %s\n", err)
	}

	if err := app.RabbitConnection.Close(); err != nil {
		log.Fatalf("faild to close rabbitmq connection: %s\n", err)
	}

	log.Println("gracefully shutting down api server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shutting down server: %s\n", err)
	}

	log.Println("server exited")
}
