package helper

import (
	"math/rand"
	"time"
)

// GenerateSlice generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size)
	rng := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 0; i < size; i++ {
		slice[i] = rng.Intn(999)
	}
	return slice
}
