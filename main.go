package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jvikstedt/blunote/pkg/db"
	"github.com/jvikstedt/blunote/pkg/handlers"
	"github.com/jvikstedt/blunote/pkg/models"
)

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

	db, err := db.NewPgDb(databaseURL)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	db.Clear()
	db.EnsureTables()
	db.Seed()

	model := models.New(db)

	handler := handlers.New(logger, model)
	http.ListenAndServe(":"+port, handler)
}
