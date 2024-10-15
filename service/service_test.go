package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{"Valid simple email", "user@example.com", true},
		{"Valid email with numbers", "user123@example.com", true},
		{"Valid email with dots", "user.name@example.com", true},
		{"Valid email with subdomain", "user@subdomain.example.com", true},
		{"Valid email with dash in domain", "user@example-domain.com", true},
		{"Valid email with uppercase", "USER@EXAMPLE.COM", true},
		{"Invalid email with IP address", "user@[192.168.0.1]", false},
		{"Invalid email without @", "userexample.com", false},
		{"Invalid email with multiple @", "user@domain@example.com", false},
		{"Invalid email with special characters", "user!#$%&'*+-/=?^_`{|}~@example.com", false},
		{"Invalid email with spaces", "user @example.com", false},
		{"Invalid email with only domain", "@example.com", false},
		{"Invalid email with plus", "user+tag@example.com", false},
		{"Invalid email with only local part", "user@", false},
		{"Invalid email with invalid TLD", "user@example.c", false},
		{"Invalid email with double dots", "user..name@example.com", false},
		{"Invalid email starting with dot", ".user@example.com", false},
		{"Invalid email ending with dot", "user.@example.com", false},
		{"Invalid email with dot before @", "user.@example.com", false},
		{"Empty string", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidEmail(tc.email)
			assert.Equal(t, tc.expected, result, "Email: %s", tc.email)
		})
	}
}

// MockFeedbackRepository is a mock type for the FeedbackRepository
type MockFeedbackRepository struct {
	mock.Mock
}

func (m *MockFeedbackRepository) SaveFeedback(name, email, feedback string) error {
	args := m.Called(name, email, feedback)
	return args.Error(0)
}

func TestSubmitFeedback(t *testing.T) {
	mockRepo := new(MockFeedbackRepository)
	service := &FeedbackService{repo: mockRepo}

	testCases := []struct {
		name          string
		input         FeedbackInput
		mockBehavior  func()
		expectedError string
	}{
		{
			name: "Valid input",
			input: FeedbackInput{
				Name:     "John Doe",
				Email:    "john@example.com",
				Feedback: "Great service!",
			},
			mockBehavior: func() {
				mockRepo.On("SaveFeedback", "John Doe", "john@example.com", "Great service!").Return(nil)
			},
			expectedError: "",
		},
		{
			name: "Empty name",
			input: FeedbackInput{
				Name:     "",
				Email:    "john@example.com",
				Feedback: "Great service!",
			},
			mockBehavior:  func() {},
			expectedError: "name cannot be empty",
		},
		{
			name: "Invalid email",
			input: FeedbackInput{
				Name:     "John Doe",
				Email:    "invalid-email",
				Feedback: "Great service!",
			},
			mockBehavior:  func() {},
			expectedError: "invalid email format",
		},
		{
			name: "Empty feedback",
			input: FeedbackInput{
				Name:     "John Doe",
				Email:    "john@example.com",
				Feedback: "",
			},
			mockBehavior:  func() {},
			expectedError: "feedback cannot be empty",
		},
		{
			name: "Repository error",
			input: FeedbackInput{
				Name:     "John Doe",
				Email:    "john@example.com",
				Feedback: "Great service!",
			},
			mockBehavior: func() {
				mockRepo.On("SaveFeedback", "John Doe", "john@example.com", "Great service!").Return(errors.New("database error"))
			},
			expectedError: "database error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mock expectations
			mockRepo.ExpectedCalls = nil

			// Set up mock behavior
			tc.mockBehavior()

			// Call the method
			err := service.SubmitFeedback(tc.input)

			// Check the result
			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			// Verify that all expected mock calls were made
			mockRepo.AssertExpectations(t)
		})
	}
}
