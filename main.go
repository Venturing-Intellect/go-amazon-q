package main

import (
	"log"
	"net/http"
	"os"

	"go-amazon-q/controller"
	"go-amazon-q/repository"
	"go-amazon-q/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize the database connection
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := repository.InitDB(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repository, service, and controller
	repo := repository.NewFeedbackRepository(db)
	svc := service.NewFeedbackService(repo)
	ctrl := controller.NewFeedbackController(svc)

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/feedback", ctrl.SubmitFeedback).Methods("POST")

	// Create a CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow requests from your React app
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap your router with the CORS handler
	handler := c.Handler(r)

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
