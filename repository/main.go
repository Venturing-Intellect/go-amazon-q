package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type FeedbackRepository struct {
	db *sql.DB
}

func NewFeedbackRepository(db *sql.DB) *FeedbackRepository {
	return &FeedbackRepository{db: db}
}

func (r *FeedbackRepository) SaveFeedback(name, email, feedback string) error {
	query := `INSERT INTO feedback (name, email, feedback) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, name, email, feedback)
	return err
}

func InitDB(user, password, host, port, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}
