package controller

import (
	"encoding/json"
	"go-amazon-q/service"
	"net/http"
)

type FeedbackService interface {
	SubmitFeedback(input service.FeedbackInput) error
}

type FeedbackController struct {
	service FeedbackService
}

func NewFeedbackController(s FeedbackService) *FeedbackController {
	return &FeedbackController{service: s}
}

func (c *FeedbackController) SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback service.FeedbackInput
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the input data
	if feedback.Name == "" || feedback.Email == "" || feedback.Feedback == "" {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	err = c.service.SubmitFeedback(feedback)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Feedback submitted successfully"})
}
