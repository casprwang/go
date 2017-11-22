package play

import (
	"math/rand"
	"time"
)

func Shuffle(arr []int) []int {
	len := len(arr)
	shuffled := make([]int, len)

	var ran int

	for i := 0; i < len; i++ {

		rand.Seed(time.Now().UnixNano())
		ran = rand.Intn(i + 1)

		if ran != i {
			shuffled[i] = shuffled[ran]
		}

		shuffled[ran] = arr[i]
	}

	return shuffled
}
