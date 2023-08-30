package ArrayGenerator

import (
	"math/rand"
)

func ArrayGenerate(size int) []int {
	var array []int
	for i := 0; i <= size; i++ {
		array = append(array, rand.Intn(10000))
	}

	return array
}
