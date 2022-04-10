package dsa

import "sort"

//first attempt
func FindMax1(in []int) int {
	max := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}
	return max
}

//second attempt fixed panic
func FindMax2(in []int) int {
	if len(in) == 0 {
		return 0
	}
	max := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}
	return max
}

//iteration with range
func FindMax3(in []int) (max int) {
	for _, v := range in {
		if v > max {
			max = v
		}
	}
	return max
}

//fixed if all numbers in slice is negative number
func FindMax4(in []int) (max int) {
	for i, v := range in {
		if v > max || i == 0 {
			max = v
		}
	}
	return max
}

//find another solution from internet but test is failed
func FindMax5(in []int) (max int) {
	for i := 0; i < len(in)-1; i++ {
		if in[i] > in[i+1] {
			if in[i] > max {
				max = in[i]
			}
		}
	}
	return max
}

//with sorting
func FindMax6(in []int) int {
	ln := len(in)
	if ln == 0 {
		return 0
	}
	sort.Ints(in)
	return in[ln-1]
}
