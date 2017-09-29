package models

import (
	"strings"

	"github.com/kennygrant/sanitize"
)

type db interface {
	GetNotes() ([]*Note, error)
	SearchNotes(string) ([]*Note, error)
	GetOneNote(string) (*Note, error)
	UpdateNote(*Note) error
	CreateNote(*Note) error
}

type Model struct {
	db db
}

func New(db db) *Model {
	return &Model{db: db}
}

func (m *Model) GetNotes() ([]*Note, error) {
	return m.db.GetNotes()
}

func (m *Model) SearchNotes(input string) ([]*Note, error) {
	input = strings.Join(strings.Split(input, " "), "|")
	return m.db.SearchNotes(input)
}

func (m *Model) GetOneNote(id string) (*Note, error) {
	return m.db.GetOneNote(id)
}

func (m *Model) UpdateNote(note *Note) error {
	note.BodyTextOnly = sanitize.HTML(note.BodyHTML)
	note.BodyTextOnly = filterNewLines(note.BodyTextOnly)
	return m.db.UpdateNote(note)
}

func (m *Model) CreateNote(note *Note) error {
	note.BodyTextOnly = sanitize.HTML(note.BodyHTML)
	note.BodyTextOnly = filterNewLines(note.BodyTextOnly)
	return m.db.CreateNote(note)
}

func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		if r == 10 {
			return ' '
		}
		return r
	}, s)
}
