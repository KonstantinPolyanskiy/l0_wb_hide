package storage

import (
	"github.com/jmoiron/sqlx"
	"l0_wb_hide/internal/models"
	"l0_wb_hide/internal/storage/order"
)

type Order interface {
	Get(id int) (models.Order, error)
	Save(order models.Order) (int, error)
}

type Repository struct {
	Order
}

func New(db *sqlx.DB) Repository {
	return Repository{
		Order: order.NewOrderRepository(db),
	}
}
