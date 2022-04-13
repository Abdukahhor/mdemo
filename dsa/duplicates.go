package dsa

import "sort"

//O(n^2)
func duplicates(in []int) []int {
	if len(in) < 2 {
		return in
	}

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i] == in[j] {
				in = append(in[:i], in[i+1:]...)
			}
		}
	}
	return in
}

//O(n^2)
func duplicates2(in []int) []int {
	if len(in) < 2 {
		return in
	}

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i] == in[j] {
				in[i] = in[len(in)-1]
				in = in[:len(in)-1]
			}
		}
	}
	return in
}

//O(n)
func duplicates3(in []int) []int {
	if len(in) < 2 {
		return in
	}
	sort.Ints(in)
	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			in = append(in[:i], in[i+1:]...)
		}
	}
	return in
}

//O(n)
func duplicates4(in []int) []int {
	if len(in) < 2 {
		return in
	}
	d := make(map[int]bool)
	var rs []int
	for _, v := range in {
		if _, ok := d[v]; !ok {
			d[v] = true
			rs = append(rs, v)
		}
	}
	return rs
}
