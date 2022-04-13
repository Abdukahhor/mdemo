package dsa

import (
	"fmt"
	"testing"
)

func TestDuplicates(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 83, 63, 70,
			37, 34, 83, 17,
			19, 97, 9, 17,
		}, res: 14},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 1},
		{a: []int{-34, 0, 67, 4, -98, 5, 23, -98}, res: 7},
	}

	for _, v := range in {
		s := duplicates(v.a)
		if len(s) != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", len(s), v.res)
		}
		fmt.Println(s)
	}
}

func TestDuplicates2(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 83, 63, 70,
			37, 34, 83, 17,
			19, 97, 9, 17,
		}, res: 14},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 1},
		{a: []int{-34, 0, 67, 4, -98, 5, 23, -98}, res: 7},
	}

	for _, v := range in {
		s := duplicates2(v.a)
		if len(s) != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", len(s), v.res)
		}
		fmt.Println(s)
	}
}

func TestDuplicates3(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 83, 63, 70,
			37, 34, 83, 17,
			19, 97, 9, 17,
		}, res: 14},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 1},
		{a: []int{-34, 0, 67, 4, -98, 5, 23, -98}, res: 7},
	}

	for _, v := range in {
		s := duplicates3(v.a)
		if len(s) != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", len(s), v.res)
		}
		fmt.Println(s)
	}
}

func TestDuplicates4(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 83, 63, 70,
			37, 34, 83, 17,
			19, 97, 9, 17,
		}, res: 14},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 1},
		{a: []int{-34, 0, 67, 4, -98, 5, 23, -98}, res: 7},
	}

	for _, v := range in {
		s := duplicates4(v.a)
		if len(s) != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", len(s), v.res)
		}
		fmt.Println(s)
	}
}
