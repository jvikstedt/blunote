package models

import "time"

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (db *DB) AllNotes() ([]*Note, error) {
	notes := []*Note{}
	err := db.Select(&notes, "SELECT * FROM notes")
	return notes, err
}
