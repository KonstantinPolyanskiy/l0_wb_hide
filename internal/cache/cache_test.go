package cache

import (
	"l0_wb_hide/external/random/rand_string"
	"testing"
)

func TestCache_Add(t *testing.T) {
	cache := New(WithCapacity(3))
	data := fillTestData()

	items := make([]Item, 3)
	errors := make([]error, 3)

	for i := 0; i < 3; i++ {
		cache.Add(i+1, data[i])
	}

	for i := 0; i < 3; i++ {
		items[i], errors[i] = cache.Get(i + 1)
	}

	for i, item := range items {
		if errors[i] != nil || item.Value.(string) != data[i] {
			t.Errorf("Ожидали: %s | получили: %s | ошибка: %v\n", data[i], item.Value.(string), errors[i])
		}
	}
}

func fillTestData() []string {
	result := make([]string, 3)

	for i := 0; i < 3; i++ {
		result[i] = rand_string.New(5)
	}

	return result
}
