package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sylvie/internal/application"
	"sylvie/internal/config"
	"sylvie/internal/router"
	"syscall"

	"github.com/labstack/echo/v5"
)

func main() {
	app := application.Bootstrap()

	r := echo.New()

	server := &http.Server{
		Addr:    config.Load().Server.PORT,
		Handler: router.RegisterRoutes(r, app),
	}

	errChan := make(chan error, 1)
	go func() {
		log.Printf("api server listening in port: %s\n", config.Load().Server.PORT)
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
