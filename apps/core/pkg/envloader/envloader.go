package envloader

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = ".env"
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}
