package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

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
	defer db.Close()

	db.Clear()
	db.EnsureTables()
	db.Seed()

	model := models.New(db)
	http.Handle("/", handlers.New(logger, model))

	server := &http.Server{Addr: ":" + port}

	go func() {
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		<-sigquit

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Printf("Unable to shut down server: %v", err)
		} else {
			logger.Println("Server stopped")
		}
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Printf("%v", err)
	} else {
		logger.Println("Server closed!")
	}
}
