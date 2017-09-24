package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func Handler(env *Env, logger *log.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Default().Handler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Get("/", env.notesIndex)

			r.Route("/{noteID}", func(r chi.Router) {
				r.Get("/", env.notesGetOne)
			})
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

func (env *Env) notesGetOne(w http.ResponseWriter, r *http.Request) {
	noteID := chi.URLParam(r, "noteID")

	note, err := env.db.GetOneNote(noteID)
	if env.checkErr(err, w, http.StatusNotFound) {
		return
	}

	json.NewEncoder(w).Encode(note)
}

func (env *Env) checkErr(err error, w http.ResponseWriter, statusCode int) bool {
	if err != nil {
		env.log.Println(err)
		w.WriteHeader(statusCode)
		return true
	}
	return false
}
