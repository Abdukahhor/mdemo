package dsa

import (
	"fmt"
	"testing"
)

func TestRandNumSlice(t *testing.T) {
	s := randNumSlice(10)
	fmt.Println(s)
}
