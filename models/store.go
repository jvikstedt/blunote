package models

type NoteStore interface {
	EnsureTables()
	Clear()
	Seed()
	AllNotes() ([]*Note, error)
	GetOneNote(string) (*Note, error)
	SearchNotes(string) ([]*Note, error)
}

type Datastore interface {
	NoteStore
}
