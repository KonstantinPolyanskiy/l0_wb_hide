package phone_number

import (
	"math/rand"
	"strconv"
)

// New возвращает случайно сгенерированный номер формата (000) 000-00-00.
func New() string {
	var n int
	number := "+"

	for i := 0; i <= 10; i++ {
		n = rand.Intn(10)
		number += strconv.Itoa(n)
	}

	return number
}
