package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jvikstedt/blunote/models"
)

type Env struct {
	db  models.Datastore
	log *log.Logger
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://blunote:blunote@localhost/blunote?sslmode=disable"
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	db, err := models.NewDB(databaseURL)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	db.Clear()
	db.EnsureTables()
	db.Seed()

	env := &Env{db: db, log: logger}

	handler := Handler(env, logger)
	http.ListenAndServe(":"+port, handler)
}
