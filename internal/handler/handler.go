package handler

import (
	"encoding/json"
	"gotinyurl/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

type request struct {
	URL string `json:"url"`
}

type response struct {
	ShortURL string `json:"short_url"`
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	short, err := h.service.Shorten(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := response{ShortURL: "http://localhost:8080/" + short}
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Path[len("/"):]
	original, err := h.service.Resolve(vars)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
}
