package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	port := os.Getenv("HTTP_PORT")
	log.Printf("Number of CPUs available:%v", runtime.NumCPU())

	// Create a server instance
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: New(),
	}

	// Set up panic recovery
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic recovery: %v", err)
		}
	}()

	// Start the server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Create a channel to wait for shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for a shutdown signal
	<-stop
	log.Print("Shutting down server...")

	// Create a context with a timeout to allow existing connections to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shut down server: %v", err)
	}

	log.Print("Server shut down successfully")
}

type WorkerPool struct {
}

type Worker struct {
}

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Business logic starts here!")
		w.Write([]byte("I am ready"))
		w.WriteHeader(200)
	})
	r.HandleFunc("/pain", func(w http.ResponseWriter, r *http.Request) {
		panic(errors.New("Test out panic recovery!"))
	})
	return r
}
