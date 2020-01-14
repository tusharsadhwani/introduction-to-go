package util

import (
	"math/rand"
	"time"
)

// RandomDigit generates a random digit between 0 and 9
func RandomDigit() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10)
}
