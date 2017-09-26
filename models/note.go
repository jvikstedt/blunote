package models

import "time"

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (db *DB) AllNotes() ([]*Note, error) {
	notes := []*Note{}
	err := db.Select(&notes, "SELECT * FROM notes")
	return notes, err
}

func (db *DB) SearchNotes(input string) ([]*Note, error) {
	input = "%" + input + "%"
	notes := []*Note{}
	err := db.Select(&notes, "SELECT * FROM notes WHERE title ILIKE $1", input)
	return notes, err
}

func (db *DB) GetOneNote(id string) (*Note, error) {
	note := &Note{}

	err := db.Get(note, "SELECT * FROM notes WHERE id=$1", id)
	return note, err
}
