package show

import (
	"github.com/go-chi/chi/v5"
)

func RegisterShowRoutes(r chi.Router, handler *Handler) {
	handler.RegisterRoutes(r)
}
