package main

import "testing"

func Test(t *testing.T) {
	var test = struct {
		arr  []int
		want int
	}{[]int{
		6331,
		2793,
		1640,
		9290,
		225,
		625,
		6195,
		2303,
		5685,
		1354,
	}, 9335}
	got := Medians(test.arr)
	if got != test.want {
		t.Errorf("Medians(%v) == %d, want %d", test.arr, got, test.want)
	} else {
		t.Logf("Medians(%v) == %d, want %d", test.arr, got, test.want)
	}
}
