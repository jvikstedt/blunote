package models

import (
	"time"

	"github.com/jmoiron/sqlx"
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

type DB struct {
	*sqlx.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) EnsureTables() {
	db.MustExec(schema)
}

func (db *DB) Seed() {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO notes (title, createdAt, updatedAt) VALUES ($1, $2, $3)", "Go", time.Now(), time.Now())
	tx.MustExec("INSERT INTO notes (title, createdAt, updatedAt) VALUES ($1, $2, $3)", "Ruby", time.Now(), time.Now())
	tx.MustExec("INSERT INTO notes (title, createdAt, updatedAt) VALUES ($1, $2, $3)", "JavaScript", time.Now(), time.Now())
	tx.Commit()
}

func (db *DB) Clear() {
	tx := db.MustBegin()
	tx.MustExec("DROP SCHEMA public CASCADE")
	tx.MustExec("CREATE SCHEMA public")
	tx.Commit()
}

func (db *DB) NoteStore() NoteStore {
	return db
}
