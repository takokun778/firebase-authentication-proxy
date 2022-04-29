package main

import (
	"log"
	"os"

	_ "github.com/takokun778/firebase-authentication-proxy/domain/model/env"
	_ "github.com/takokun778/firebase-authentication-proxy/domain/model/key"
)

func main() {
	log.Printf("%s server starting...\n", os.Getenv("ENV"))

	server := InitializeHTTPServer()

	server.Run()
}
