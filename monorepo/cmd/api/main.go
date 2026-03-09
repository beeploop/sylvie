package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sylvie/internal/config"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleIndex)

	server := &http.Server{
		Addr:    config.Load().PORT,
		Handler: mux,
	}

	errChan := make(chan error, 1)
	go func() {
		log.Printf("api server listening in port: %s\n", config.Load().PORT)
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

	log.Println("gracefully shutting down api server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shutting down server: %s\n", err)
	}

	log.Println("server exited")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"message"`
		Time    string `json:"time"`
	}{
		Message: "hello world",
		Time:    time.Now().Format(time.RFC3339),
	}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
