package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	dev = "dev"
)

func init() {
	if os.Getenv("ENV") == dev {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("error loading .env file %v", err)
		}
	}
}

func IsDev() bool {
	return os.Getenv("ENV") == dev
}
