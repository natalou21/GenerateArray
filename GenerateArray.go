package Generator

import "math/rand"

func ArrayGenerate(size int) []int {
	var tuple = make([]int, size)

	for i := 0; i < size; i++ {
		tuple[i] = rand.Intn(10000)
	}

	return tuple
}
