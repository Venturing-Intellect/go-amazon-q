package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSaveFeedback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewFeedbackRepository(db)

	t.Run("successful insert", func(t *testing.T) {
		name := "John Doe"
		email := "john@example.com"
		feedback := "This is a test feedback"

		mock.ExpectExec("INSERT INTO feedback (.+)").
			WithArgs(name, email, feedback).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.SaveFeedback(name, email, feedback)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		name := "John Doe"
		email := "john@example.com"
		feedback := "This is a test feedback"

		mock.ExpectExec("INSERT INTO feedback (.+)").
			WithArgs(name, email, feedback).
			WillReturnError(errors.New("database error"))

		err := repo.SaveFeedback(name, email, feedback)

		assert.EqualError(t, err, "database error")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
