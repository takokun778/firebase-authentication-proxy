package main

import (
	"log"
	"os"

	_ "firebase-authentication/driver/env"
	"firebase-authentication/driver/server"
)

func main() {
	log.Printf("%s server starting...\n", os.Getenv("ENV"))
	log.Fatal(server.Serve(os.Getenv("PORT")))
}
