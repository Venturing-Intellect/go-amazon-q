package service

import (
	"errors"
	"regexp"
)

type FeedbackRepository interface {
	SaveFeedback(name, email, feedback string) error
}

type FeedbackService struct {
	repo FeedbackRepository
}

func NewFeedbackService(r FeedbackRepository) *FeedbackService {
	return &FeedbackService{repo: r}
}

type FeedbackInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Feedback string `json:"feedback"`
}

func (s *FeedbackService) SubmitFeedback(input FeedbackInput) error {
	if input.Name == "" {
		return errors.New("name cannot be empty")
	}

	if !isValidEmail(input.Email) {
		return errors.New("invalid email format")
	}

	if input.Feedback == "" {
		return errors.New("feedback cannot be empty")
	}

	return s.repo.SaveFeedback(input.Name, input.Email, input.Feedback)
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
