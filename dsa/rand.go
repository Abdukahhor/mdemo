package dsa

import (
	"math/rand"
	"time"
)

func randNumSlice(ln int) []int {
	s := make([]int, ln)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < ln; i++ {
		n := rand.Intn(2000000) - 1000000
		s[i] = n
	}
	return s
}
