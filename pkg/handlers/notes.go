package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jvikstedt/blunote/pkg/models"
)

type NoteRequest struct {
	*models.Note

	ProtectedID        int    `json:"id"`
	ProtectedCreatedAt string `json:"created_at"`
	ProtectedUpdatedAt string `json:"updated_at"`
}

func (n *NoteRequest) Bind(r *http.Request) error {
	return nil
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

func (h handler) notesCreate(w http.ResponseWriter, r *http.Request) {
	note := &models.Note{}
	data := &NoteRequest{Note: note}
	if err := render.Bind(r, data); h.checkErr(err, w, http.StatusBadRequest) {
		return
	}

	if err := h.model.CreateNote(data.Note); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(note)
}

func (h handler) notesUpdate(w http.ResponseWriter, r *http.Request) {
	noteID := chi.URLParam(r, "noteID")

	note, err := h.model.GetOneNote(noteID)
	if h.checkErr(err, w, http.StatusNotFound) {
		return
	}

	data := &NoteRequest{Note: note}
	if err := render.Bind(r, data); h.checkErr(err, w, http.StatusBadRequest) {
		return
	}

	if err := h.model.UpdateNote(data.Note); h.checkErr(err, w, http.StatusUnprocessableEntity) {
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
