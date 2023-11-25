package main

import (
	"database/sql"
	"fmt"
	"github.com/DayDream-21/rssagg/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

const envFileName string = ".env"
const portKey string = "PORT"
const dbUrlKey string = "DB_URL"

type apiConfig struct {
	DB *database.Queries
}

func main() {
	loadError := godotenv.Load(envFileName)
	if loadError != nil {
		log.Fatalf("Error loading %s file\n", envFileName)
	}

	port := os.Getenv(portKey)
	if port == "" {
		log.Fatalf("'%s' environment variable is not set\n", portKey)
	}

	dbUrl := os.Getenv(dbUrlKey)
	if dbUrl == "" {
		log.Fatalf("'%s' environment variable is not set\n", dbUrlKey)
	}

	conn, connectionError := sql.Open("postgres", dbUrl)
	if connectionError != nil {
		log.Fatalf("Error connecting to database: %s\n", connectionError)
	}

	apiConfig := apiConfig{
		DB: database.New(conn),
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
	v1Router.Post("/users", apiConfig.handlerCreateUser)

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
