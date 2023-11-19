package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

const envFileName string = ".env"
const portKey string = "PORT"

func main() {
	loadError := godotenv.Load(envFileName)
	if loadError != nil {
		log.Fatalf("Error loading %s file\n", envFileName)
	}

	port := os.Getenv(portKey)
	if port == "" {
		log.Fatalf("'%s' environment variable is not set\n", portKey)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Starting server on port: %s\n", port)
	serverStartError := server.ListenAndServe()
	if serverStartError != nil {
		log.Fatalf("Error starting server: %s\n", serverStartError)
	}

	fmt.Printf("Port: %s\n", port)
}
