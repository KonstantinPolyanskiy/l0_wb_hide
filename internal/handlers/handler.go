package handlers

import (
	"github.com/go-chi/chi/v5"
	"l0_wb_hide/internal/broker/stream"
	"l0_wb_hide/internal/handlers/order"
	"l0_wb_hide/internal/usecase"
)

type Handler struct {
	order  order.Order
	stream stream.Stream
}

func New(service usecase.Service, stream stream.Stream) Handler {
	return Handler{
		order: order.NewHandler(service.Order, stream),
	}
}

func (h Handler) Init() *chi.Mux {
	r := chi.NewRouter()

	h.order.ProcessMessage()

	r.Get("/order?id={id}", h.order.Get())

	return r
}
