package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

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
