package main

import (
	"context"
	"errors"
	"log"
	http2 "net/http"
	"os/signal"
	"syscall"
	"time"

	"wb-tech/l2/18/internal/config"
	"wb-tech/l2/18/internal/http"
	"wb-tech/l2/18/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	conf := config.NewConfig()

	serv := service.NewService()
	handler := http.NewHandler(serv)
	mux := http.NewMux(handler)

	logged := http.LoggingMiddleware(mux)

	server := &http2.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: logged,
	}

	go func() {
		log.Println("Starting server on port:" + conf.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(http2.ErrServerClosed, err) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to gracefully shutdown server: %v", err)
	}
	log.Println("Server gracefully stopped")
}
