package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFeedbackRepository struct {
	mock.Mock
}

func (m *MockFeedbackRepository) SaveFeedback(name, email, feedback string) error {
	args := m.Called(name, email, feedback)
	return args.Error(0)
}
func TestSubmitFeedback(t *testing.T) {
	mockRepo := new(MockFeedbackRepository)
	service := NewFeedbackService(mockRepo)

	t.Run("valid input", func(t *testing.T) {
		input := FeedbackInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Feedback: "This is a test feedback",
		}
		mockRepo.On("SaveFeedback", input.Name, input.Email, input.Feedback).Return(nil)

		err := service.SubmitFeedback(input)

		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("empty name", func(t *testing.T) {
		input := FeedbackInput{
			Name:     "", // Empty name
			Email:    "john@example.com",
			Feedback: "This is a test feedback",
		}

		err := service.SubmitFeedback(input)

		assert.EqualError(t, err, "name cannot be empty")
	})

	t.Run("invalid email", func(t *testing.T) {
		input := FeedbackInput{
			Name:     "John Doe",
			Email:    "invalid-email.com",
			Feedback: "This is a test feedback",
		}

		err := service.SubmitFeedback(input)

		assert.EqualError(t, err, "invalid email format")
	})

	t.Run("empty feedback", func(t *testing.T) {
		input := FeedbackInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Feedback: "",
		}

		err := service.SubmitFeedback(input)

		assert.EqualError(t, err, "feedback cannot be empty")
	})

}
