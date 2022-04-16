package main

import (
	"log"
	"os"

	_ "github.com/takokun778/firebase-authentication-proxy/driver/env"
	"github.com/takokun778/firebase-authentication-proxy/driver/server"
)

func main() {
	log.Printf("%s server starting...\n", os.Getenv("ENV"))
	log.Fatal(server.Serve(os.Getenv("PORT")))
}
