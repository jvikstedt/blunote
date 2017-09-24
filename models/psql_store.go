package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var tables = []string{
	"notes",
}

var schema = `
CREATE TABLE IF NOT EXISTS notes (
	id serial NOT NULL,
	title character varying(100) NOT NULL,
	body text,
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
	db.Clear()

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO notes (title, body, createdAt, updatedAt) VALUES ($1, $2, $3, $4)", "Go", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", time.Now(), time.Now())
	tx.MustExec("INSERT INTO notes (title, body, createdAt, updatedAt) VALUES ($1, $2, $3, $4)", "Ruby", "Aliquam ex mauris, posuere in mattis et, dignissim aliquet elit.", time.Now(), time.Now())
	tx.MustExec("INSERT INTO notes (title, body, createdAt, updatedAt) VALUES ($1, $2, $3, $4)", "JavaScript", "Donec a augue congue, pellentesque odio in, fermentum augue.", time.Now(), time.Now())
	tx.Commit()
}

func (db *DB) Clear() {
	tx := db.MustBegin()
	for _, v := range tables {
		tx.MustExec(fmt.Sprintf("DROP TABLE IF EXISTS %s", v))
	}
	tx.Commit()

	db.EnsureTables()
}

func (db *DB) NoteStore() NoteStore {
	return db
}
