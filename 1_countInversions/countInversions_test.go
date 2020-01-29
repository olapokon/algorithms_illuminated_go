package main

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{}, 0},
		{[]int{8, 4}, 1},
		{[]int{1, 6, 8, 16, 4}, 3},
	}
	for _, a := range tests {
		_, got := MergeSort(a.arr)
		if got != a.want {
			t.Errorf("MergeSort(%v) == %d, want %d", a.arr, got, a.want)
		}
	}
}
