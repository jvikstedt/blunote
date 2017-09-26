package models

type db interface {
	GetNotes() ([]*Note, error)
	SearchNotes(string) ([]*Note, error)
	GetOneNote(string) (*Note, error)
}

type Model struct {
	db
}

func New(db db) *Model {
	return &Model{db: db}
}
