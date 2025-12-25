package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/abdoulrl2028-cloud-Dev/api-golang-crud/internal/db"
	"github.com/abdoulrl2028-cloud-Dev/api-golang-crud/internal/handler"
	"github.com/abdoulrl2028-cloud-Dev/api-golang-crud/internal/repository"
	"github.com/abdoulrl2028-cloud-Dev/api-golang-crud/internal/service"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database connection
	database, err := db.GetConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(database)

	// Initialize repositories
	userRepo := repository.NewUserRepository(database)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup routes
	router := mux.NewRouter()

	// Register user routes
	userHandler.RegisterRoutes(router)

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	// Get API port from environment
	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
