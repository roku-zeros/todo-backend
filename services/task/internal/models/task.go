package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"due_date,omitempty"`
	Completed   bool      `json:"is_completed,omitempty"`
}
