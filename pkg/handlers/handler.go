package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jvikstedt/blunote/pkg/models"
)

func New(logger *log.Logger, m *models.Model) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Default().Handler)

	h := handler{logger, m}

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Get("/", h.notesIndex)

			r.Route("/{noteID}", func(r chi.Router) {
				r.Get("/", h.notesGetOne)
			})
		})
	})
	return r
}

type handler struct {
	log   *log.Logger
	model *models.Model
}

func (h handler) notesIndex(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	notes, err := h.model.SearchNotes(search)
	if h.checkErr(err, w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(notes)
}

func (h handler) notesGetOne(w http.ResponseWriter, r *http.Request) {
	noteID := chi.URLParam(r, "noteID")

	note, err := h.model.GetOneNote(noteID)
	if h.checkErr(err, w, http.StatusNotFound) {
		return
	}

	json.NewEncoder(w).Encode(note)
}

func (h handler) checkErr(err error, w http.ResponseWriter, statusCode int) bool {
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(statusCode)
		return true
	}
	return false
}
