package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-amazon-q/controller"
	"go-amazon-q/repository"
	"go-amazon-q/service"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	// Setup
	db, err := setupTestDatabase()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewFeedbackRepository(db)
	svc := service.NewFeedbackService(repo)
	ctrl := controller.NewFeedbackController(svc)

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(ctrl.SubmitFeedback))
	defer server.Close()

	// Test case
	t.Run("submit feedback", func(t *testing.T) {
		input := service.FeedbackInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Feedback: "This is an integration test feedback",
		}

		body, _ := json.Marshal(input)
		resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(body))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		// Verify the feedback was saved in the database
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM feedback WHERE name = $1 AND email = $2 AND feedback = $3",
			input.Name, input.Email, input.Feedback).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 1, count, "Expected one feedback entry to be saved")
	})
}

func setupTestDatabase() (*sql.DB, error) {
	// Use environment variables or a test configuration file for these values
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	db, err := repository.InitDB(user, password, host, port, dbname)
	if err != nil {
		return nil, err
	}

	// Clear the feedback table before running tests
	_, err = db.Exec("DELETE FROM feedback")
	if err != nil {
		return nil, err
	}

	return db, nil
}
