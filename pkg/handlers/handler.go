package handlers

import (
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
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	r.Use(cors.Handler)

	h := handler{logger, m}

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Get("/", h.notesIndex)
			r.Post("/", h.notesCreate)

			r.Route("/{noteID}", func(r chi.Router) {
				r.Get("/", h.notesGetOne)
				r.Put("/", h.notesUpdate)
			})
		})
	})
	return r
}

type handler struct {
	log   *log.Logger
	model *models.Model
}
