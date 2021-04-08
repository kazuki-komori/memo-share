package entity

import "time"

type Content struct {
	ID           string    `json:"id"`
	ContentTitle string    `json:"content_title"`
	Description  string    `json:"description"`
	UserID       string    `json:"user_id"`
	SubjectID    string    `json:"subject_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
