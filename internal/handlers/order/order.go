package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"l0_wb_hide/internal/usecase"
	"net/http"
	"strconv"
)

type Order struct {
	service usecase.Order
}

func NewHandler(service usecase.Order) Order {
	return Order{
		service: service,
	}
}

func (h Order) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if idStr == "" {
			http.Error(w, "empty id", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id is not a integer", http.StatusBadRequest)
			return
		}

		order, err := h.service.Get(id)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"result": order,
		})
	}
}
