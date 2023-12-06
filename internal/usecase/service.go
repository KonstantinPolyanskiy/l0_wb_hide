package usecase

import (
	"l0_wb_hide/internal/models"
	"l0_wb_hide/internal/storage"
	"l0_wb_hide/internal/usecase/order"
)

type Order interface {
	Get(id int) (models.Order, error)
	Save(order models.Order) error
}
type Service struct {
	Order
}

func New(repo storage.Repository, cache order.Cache) Service {
	return Service{
		Order: order.NewOrderService(repo.Order, cache),
	}
}
