package models

import "time"

type Note struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	BodyHTML     string    `json:"body_html" db:"body_html"`
	BodyTextOnly string    `json:"body_text_only" db:"body_text_only"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
