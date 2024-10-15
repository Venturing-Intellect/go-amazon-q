package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-amazon-q/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFeedbackService struct {
	mock.Mock
}

func (m *MockFeedbackService) SubmitFeedback(input service.FeedbackInput) error {
	args := m.Called(input)
	return args.Error(0)
}

func TestSubmitFeedback(t *testing.T) {
	mockService := new(MockFeedbackService)
	controller := NewFeedbackController(mockService)

	t.Run("valid input", func(t *testing.T) {
		input := service.FeedbackInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Feedback: "This is a test feedback",
		}
		mockService.On("SubmitFeedback", input).Return(nil)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		controller.SubmitFeedback(rr, req)

		mockService.AssertExpectations(t)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("invalid request body", func(t *testing.T) {
		body := []byte(`{"invalid": "data"}`)
		req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		controller.SubmitFeedback(rr, req)

		// The service's SubmitFeedback method should not be called for invalid input
		mockService.AssertNotCalled(t, "SubmitFeedback")
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("service error", func(t *testing.T) {
		input := service.FeedbackInput{
			Name:     "John Doe",
			Email:    "invalid-email",
			Feedback: "This is a test feedback",
		}
		mockService.On("SubmitFeedback", input).Return(errors.New("invalid email format"))

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		controller.SubmitFeedback(rr, req)

		mockService.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
