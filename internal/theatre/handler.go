package theatre

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
	r.Route("/theatres", func(r chi.Router) {
		r.Get("/", h.GetAllTheatres)
		r.Get("/{id}", h.GetTheatreByID)
		r.Post("/", h.CreateTheatre)
		r.Put("/{id}", h.UpdateTheatre)
		r.Delete("/{id}", h.DeleteTheatre)
	})
}

func (h *Handler) GetAllTheatres(w http.ResponseWriter, r *http.Request) {
	theatres, err := h.service.GetAllTheatres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if theatres == nil {
		theatres = []Theatre{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(theatres)
}

func (h *Handler) GetTheatreByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	theatre, err := h.service.GetTheatreByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(theatre)
}

func (h *Handler) CreateTheatre(w http.ResponseWriter, r *http.Request) {
	var theatre Theatre
	if err := json.NewDecoder(r.Body).Decode(&theatre); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	createdTheatre, err := h.service.CreateTheatre(theatre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTheatre)
}

func (h *Handler) UpdateTheatre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var theatre Theatre
	if err := json.NewDecoder(r.Body).Decode(&theatre); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	updatedTheatre, err := h.service.UpdateTheatre(id, theatre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTheatre)
}

func (h *Handler) DeleteTheatre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteTheatre(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
