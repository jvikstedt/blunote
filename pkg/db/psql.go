package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/blunote/pkg/models"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS notes (
	id serial NOT NULL,
	title character varying(100) NOT NULL,
	createdAt timestamptz,
	updatedAt timestamptz default current_timestamp,
	CONSTRAINT notes_pkey PRIMARY KEY (id)
);`

type pgDb struct {
	*sqlx.DB
}

func NewPgDb(dataSourceName string) (*pgDb, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &pgDb{db}, nil
}

func (db *pgDb) EnsureTables() {
	db.MustExec(schema)
}

func (db *pgDb) Seed() {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO notes (title, createdAt, updatedAt) VALUES ($1, $2, $3)", "Go", time.Now(), time.Now())
	tx.MustExec("INSERT INTO notes (title, createdAt, updatedAt) VALUES ($1, $2, $3)", "Ruby", time.Now(), time.Now())
	tx.MustExec("INSERT INTO notes (title, createdAt, updatedAt) VALUES ($1, $2, $3)", "JavaScript", time.Now(), time.Now())
	tx.Commit()
}

func (db *pgDb) Clear() {
	tx := db.MustBegin()
	tx.MustExec("DROP SCHEMA public CASCADE")
	tx.MustExec("CREATE SCHEMA public")
	tx.Commit()
}

func (db *pgDb) GetNotes() ([]*models.Note, error) {
	notes := []*models.Note{}
	err := db.Select(&notes, "SELECT * FROM notes")
	return notes, err
}

func (db *pgDb) SearchNotes(input string) ([]*models.Note, error) {
	input = "%" + input + "%"
	notes := []*models.Note{}
	err := db.Select(&notes, "SELECT * FROM notes WHERE title ILIKE $1", input)
	return notes, err
}

func (db *pgDb) GetOneNote(id string) (*models.Note, error) {
	note := &models.Note{}

	err := db.Get(note, "SELECT * FROM notes WHERE id=$1", id)
	return note, err
}
