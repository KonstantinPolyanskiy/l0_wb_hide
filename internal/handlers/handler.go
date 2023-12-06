package handlers

import (
	"github.com/go-chi/chi/v5"
	"l0_wb_hide/internal/handlers/order"
	"l0_wb_hide/internal/usecase"
)

type Handler struct {
	order order.Order
}

func New(service usecase.Service) Handler {
	return Handler{
		order: order.NewHandler(service.Order),
	}
}

func (h Handler) Init() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/order", func(r chi.Router) {
		r.Get("{id}", h.order.Get())
	})

	return r
}
