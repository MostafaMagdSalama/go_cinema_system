package show

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/shows", func(r chi.Router) {
		r.Get("/", h.GetAllShows)
		r.Get("/{id}", h.GetShowByID)
		r.Post("/", h.CreateShow)
		r.Put("/{id}", h.UpdateShow)
		r.Delete("/{id}", h.DeleteShow)
	})
}

func (h *Handler) GetAllShows(w http.ResponseWriter, r *http.Request) {
	shows, err := h.service.GetAllShows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(shows)
}

func (h *Handler) GetShowByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	show, err := h.service.GetShowByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(show)
}

func (h *Handler) CreateShow(w http.ResponseWriter, r *http.Request) {
	var show Show
	if err := json.NewDecoder(r.Body).Decode(&show); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	createdShow, err := h.service.CreateShow(show)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdShow)
}

func (h *Handler) UpdateShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var show Show
	if err := json.NewDecoder(r.Body).Decode(&show); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	updatedShow, err := h.service.UpdateShow(id, show)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedShow)
}

func (h *Handler) DeleteShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteShow(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
