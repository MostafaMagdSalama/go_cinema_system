package seat

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
	r.Route("/seats", func(r chi.Router) {
		r.Get("/", h.GetAllSeats)
		r.Get("/{id}", h.GetSeatByID)
		r.Get("/theatre/{theatreId}", h.GetSeatsByTheatre)
		r.Post("/", h.CreateSeat)
		r.Put("/{id}", h.UpdateSeat)
		r.Delete("/{id}", h.DeleteSeat)
	})
}

func (h *Handler) GetAllSeats(w http.ResponseWriter, r *http.Request) {
	seats, err := h.service.GetAllSeats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if seats == nil {
		seats = []Seat{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seats)
}

func (h *Handler) GetSeatByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	seat, err := h.service.GetSeatByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(seat)
}

func (h *Handler) GetSeatsByTheatre(w http.ResponseWriter, r *http.Request) {
	theatreID, err := strconv.Atoi(chi.URLParam(r, "theatreId"))
	if err != nil {
		http.Error(w, "Invalid Theatre ID", http.StatusBadRequest)
		return
	}
	seats, err := h.service.GetSeatsByTheatreID(theatreID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(seats)
}

func (h *Handler) CreateSeat(w http.ResponseWriter, r *http.Request) {
	var seat Seat
	if err := json.NewDecoder(r.Body).Decode(&seat); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	createdSeat, err := h.service.CreateSeat(seat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdSeat)
}

func (h *Handler) UpdateSeat(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var seat Seat
	if err := json.NewDecoder(r.Body).Decode(&seat); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	updatedSeat, err := h.service.UpdateSeat(id, seat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedSeat)
}

func (h *Handler) DeleteSeat(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteSeat(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
