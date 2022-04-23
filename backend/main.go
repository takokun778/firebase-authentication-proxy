package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/takokun778/firebase-authentication-proxy/driver/env"
	"github.com/takokun778/firebase-authentication-proxy/driver/server"
)

const shutdownTime = 10

func main() {
	log.Printf("%s server starting...\n", os.Getenv("ENV"))

	srv := server.Initialize(os.Getenv("PORT"))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("Server closed with error: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
	}

	log.Println("Server shutdown")
}
