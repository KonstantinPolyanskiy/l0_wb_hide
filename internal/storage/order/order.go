package order

import (
	"github.com/jmoiron/sqlx"
	"l0_wb_hide/internal/models"
	"math/rand"
)

type Order struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) Order {
	return Order{
		db: db,
	}
}

func (o Order) Get(id int) (models.Order, error) {
	return models.Order{}, nil
}
func (o Order) Save(savedOrder models.Order) (int, error) {
	return rand.Intn(100), nil
}
