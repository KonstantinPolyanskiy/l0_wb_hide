package rand_string

import "math/rand"

const ascii = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// New возвращает строку длинной n из латинских букв.
func New(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = ascii[rand.Intn(len(ascii))]
	}

	return string(b)
}
