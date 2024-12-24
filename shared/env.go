package shared

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVar(key string) string {
	val := os.Getenv(key)

	if val != "" {
		return val
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
