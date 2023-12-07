package order

import (
	"l0_wb_hide/internal/cache"
	"l0_wb_hide/internal/models"
	"l0_wb_hide/internal/storage"
	"log"
)

type Cache interface {
	Add(key int, value interface{})
	Get(key int) (cache.Item, error)
}

type Order struct {
	repository storage.Order
	cache      Cache
}

func NewOrderService(repository storage.Order, cache Cache) Order {
	return Order{
		repository: repository,
		cache:      cache,
	}
}

func (o Order) Get(id int) (models.Order, error) {
	value, err := o.cache.Get(id)
	if err != nil {
		log.Printf("Ошибка получения элемента из кэша - %s\v", err)
	}

	orderResult, ok := value.Value.(models.Order)
	if !ok {
		log.Printf("В кеше лежит не models.Order")
	}

	orderResult, err = o.repository.Get(id)
	if err != nil {
		return models.Order{}, err
	}

	return orderResult, nil
}
func (o Order) Save(savedOrder models.Order) error {
	id, err := o.repository.Save(savedOrder)
	if err != nil {
		return err
	}
	log.Printf("запись с id %d сохранена", id)

	o.cache.Add(id, savedOrder)

	return nil
}
