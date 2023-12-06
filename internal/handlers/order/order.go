package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"l0_wb_hide/internal/models"
	"l0_wb_hide/internal/usecase"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Order struct {
	service usecase.Order
	Recipienter
}

type Recipienter interface {
	TakeOrder() (models.Order, error)
}

func NewHandler(service usecase.Order, rec Recipienter) Order {
	return Order{
		service:     service,
		Recipienter: rec,
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
func (h Order) ProcessMessage() {
	orderChan := make(chan models.Order)
	errChan := make(chan error)

	go func() {
		for {
			order, err := h.TakeOrder()
			if err != nil {
				writeToErrChan(errChan, err)
				return
			} else {
				writeToOrderChan(orderChan, order)
				log.Println("Записано")
			}

			// Проверяем каждые 500 микросекунд, есть ли непрочитанный Order
			time.Sleep(500 * time.Microsecond)
		}
	}()

	go func() {
		for {
			select {
			case order := <-orderChan:
				err := h.service.Save(order)
				if err != nil {
					writeToErrChan(errChan, err)
				}
			case err := <-errChan:
				log.Printf("Ошибка в обработке ордера со стрима - %s\n", err)
			}
		}
	}()

	time.Sleep(500 * time.Microsecond)
}

func writeToOrderChan(orderChan chan<- models.Order, order models.Order) {
	orderChan <- order
}
func writeToErrChan(errChan chan<- error, err error) {
	errChan <- err
}
