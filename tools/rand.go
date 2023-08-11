package tools

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	var letters = []rune("bcdfghijklmnpqrstwxyz")
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
