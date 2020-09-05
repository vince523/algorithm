package algorithm

import (
	"math/rand"
	"time"
)

func GenRandomArray(length, max int) []int {
	arr := make([]int, 0)
	for i := 0; i < length; i++ {
		seed := time.Now().UnixNano()
		r := rand.New(rand.NewSource(seed))
		num := r.Intn(max)
		arr = append(arr, num)
	}

	return arr
}
