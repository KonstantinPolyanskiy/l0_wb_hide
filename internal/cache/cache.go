package cache

import (
	"errors"
	"sync"
	"time"
)

var NoExistItem = errors.New("item no exist in cache")

type Cache struct {
	// Capacity - сколько всего элементов может содержаться в кеше
	Capacity int
	mu       sync.RWMutex
	Storage  map[int]Item
}

type Item struct {
	Value      interface{}
	addingTime time.Time
}

type Option func(cache *Cache)

func New(opts ...Option) *Cache {
	cache := &Cache{
		Capacity: 0,
		mu:       sync.RWMutex{},
		Storage:  make(map[int]Item),
	}

	for _, opt := range opts {
		opt(cache)
	}

	return cache
}

func (c *Cache) Add(key int, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.elemsNow() >= c.Capacity {
		c.mu.Unlock()
		c.deleteOldest()
		c.mu.Lock()
	}

	item := Item{
		Value:      value,
		addingTime: time.Now(),
	}

	c.Storage[key] = item
}
func (c *Cache) Get(key int) (Item, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.Storage[key]
	if !ok {
		return Item{}, NoExistItem
	}

	return item, nil
}

// WithCapacity задает максимальную вместимость хранилища.
// По умолчанию - 50 элементов.
func WithCapacity(cap int) Option {
	return func(c *Cache) {
		c.Capacity = cap
	}
}

// elemsNow возвращает количество элементов, находящихся в кеше в данный момент
func (c *Cache) elemsNow() int {
	return len(c.Storage)
}

// deleteOldest удаляет элемент, который находится в кеше дольше всего относительно time.Now().
func (c *Cache) deleteOldest() {
	c.mu.Lock()
	defer c.mu.Unlock()

	oldestKey := -1
	oldestTime := time.Now()

	for key, item := range c.Storage {
		if item.addingTime.Before(oldestTime) {
			oldestKey = key
			oldestTime = item.addingTime
		}
	}

	if oldestKey >= 1 {
		delete(c.Storage, oldestKey)
	}
}
