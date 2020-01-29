package main

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{}, 0},
		{[]int{8, 4}, 1},
		{[]int{8, 12, 4}, 2},
		{[]int{8, 6, 1, 16, 4}, 6},
		{[]int{2148, 9058, 7742, 3153, 6324, 609, 7628, 5469, 7017, 50}, 21},
	}
	for _, a := range tests {
		got := QuickSort(a.arr, 0, len(a.arr)-1)
		if got != a.want {
			t.Errorf("QuickSort(%v) == %d comparisons, want %d", a.arr, got, a.want)
		}
	}
}
