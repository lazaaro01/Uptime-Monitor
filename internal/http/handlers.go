package http

import (
	"encoding/json"
	"net/http"

	"uptime-monitor/internal/check"
	"uptime-monitor/internal/storage"
)

type Handler struct {
	store *storage.MemoryStorage
}

func NewHandler(store *storage.MemoryStorage) *Handler {
	return &Handler{store: store}
}

func (h *Handler) CreateCheck(w http.ResponseWriter, r *http.Request) {
	var input struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	c := check.Check{
		ID:     generateID(),
		URL:    input.URL,
		Status: check.StatusDown,
	}

	h.store.Save(c)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *Handler) ListChecks(w http.ResponseWriter, r *http.Request) {
	checks := h.store.GetAll()
	json.NewEncoder(w).Encode(checks)
}