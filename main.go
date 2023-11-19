package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envFileName string = ".env"
const portKey string = "PORT"

func main() {
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Fatalf("Error loading %s file\n", envFileName)
	}

	port := os.Getenv(portKey)
	if port == "" {
		log.Fatalf("'%s' environment variable is not set\n", portKey)
	}

	fmt.Printf("Port: %s\n", port)
}
