package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleIndex)

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	go func() {
		log.Println("api server listening in port: 3000")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to run api server: %s\n", err)
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT)
	<-quitChan

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
