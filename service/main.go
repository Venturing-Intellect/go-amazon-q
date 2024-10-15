package service

import (
	"errors"
	"regexp"
	"strings"
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
	// More restrictive regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return false
	}

	// Additional checks
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	local, domain := parts[0], parts[1]

	// Check local part
	if strings.HasPrefix(local, ".") || strings.HasSuffix(local, ".") || strings.Contains(local, "..") {
		return false
	}

	// Disallow '+' in local part
	if strings.Contains(local, "+") {
		return false
	}

	// Check domain part
	if strings.HasPrefix(domain, "[") && strings.HasSuffix(domain, "]") {
		return false // Disallow IP addresses in domain
	}

	// Ensure domain has at least one dot and the last part is at least 2 characters
	domainParts := strings.Split(domain, ".")
	if len(domainParts) < 2 || len(domainParts[len(domainParts)-1]) < 2 {
		return false
	}

	return true
}
