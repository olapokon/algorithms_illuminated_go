package main

import "testing"

func Test(t *testing.T) {
	var test = struct {
		m    map[int]int
		want int
	}{map[int]int{
		-3: -3,
		-1: -1,
		1:  1,
		2:  2,
		6:  6,
		7:  7,
		9:  9,
		11: 11,
	}, 8}
	c := make(chan int)
	go TwoSUM(3, 10, test.m, c)
	got := <-c
	if got != test.want {
		t.Errorf("TwoSUM(%v) == %d, want %d", test.m, got, test.want)
	} else {
		t.Logf("TwoSUM(%v) == %d, want %d", test.m, got, test.want)
	}
}
