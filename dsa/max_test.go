package dsa

import (
	"testing"
)

func TestFindMax1(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}, res: 97},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 34},
	}

	for _, v := range in {
		max := FindMax1(v.a)
		if max != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", max, v.res)
		}
	}
}

func TestFindMax2(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}, res: 97},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 34},
	}

	for _, v := range in {
		max := FindMax2(v.a)
		if max != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", max, v.res)
		}
	}
}

func TestFindMax3(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}, res: 97},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 34},
		{a: []int{-34, 0, 67, 4, -98, 5, 23}, res: 67},
		{a: []int{-34, -7, -66, -55}, res: -7},
	}

	for _, v := range in {
		max := FindMax3(v.a)
		if max != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", max, v.res)
		}
	}
}

func TestFindMax4(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}, res: 97},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 34},
		{a: []int{-34, 0, 67, 4, -98, 5, 23}, res: 67},
		{a: []int{-34, -7, -66, -55}, res: -7},
	}

	for _, v := range in {
		max := FindMax4(v.a)
		if max != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", max, v.res)
		}
	}
}

func TestFindMax5(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}, res: 97},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 34},
		{a: []int{-34, 0, 67, 4, -98, 5, 23}, res: 67},
		{a: []int{-34, -7, -66, -55}, res: -7},
	}

	for _, v := range in {
		max := FindMax5(v.a)
		if max != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", max, v.res)
		}
	}
}

func TestFindMax6(t *testing.T) {

	in := []struct {
		a   []int
		res int
	}{
		{a: []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}, res: 97},
		{a: []int{}, res: 0},
		{a: []int{34}, res: 34},
		{a: []int{-34, 0, 67, 4, -98, 5, 23}, res: 67},
		{a: []int{-34, -7, -66, -55}, res: -7},
	}

	for _, v := range in {
		max := FindMax6(v.a)
		if max != v.res {
			t.Errorf("Max was incorrect, got: %d, want: %d.", max, v.res)
		}
	}
}
