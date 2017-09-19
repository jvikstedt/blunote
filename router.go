package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Handler(env *Env, logger *log.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Get("/", env.notesIndex)
		})
	})
	return r
}

func (env *Env) notesIndex(w http.ResponseWriter, r *http.Request) {
	notes, err := env.db.AllNotes()
	if env.checkErr(err, w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(notes)
}

func (env *Env) checkErr(err error, w http.ResponseWriter, statusCode int) bool {
	if err != nil {
		env.log.Println(err)
		w.WriteHeader(statusCode)
		return true
	}
	return false
}
