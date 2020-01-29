package main

import (
	"testing"
)

func Test(t *testing.T) {
	var test = struct {
		edges    map[int][]int
		edgesRev map[int][]int
		vertices map[int]*vertex
		want     []int
	}{
		map[int][]int{
			1: []int{2},
			2: []int{3},
			3: []int{1, 4},
			5: []int{4},
			6: []int{4, 7},
			7: []int{8},
			8: []int{6},
		},
		map[int][]int{
			1: []int{3},
			2: []int{1},
			3: []int{2},
			4: []int{3, 5, 6},
			6: []int{8},
			7: []int{6},
			8: []int{7},
		},
		map[int]*vertex{
			1: &vertex{0, false, 0},
			2: &vertex{0, false, 0},
			3: &vertex{0, false, 0},
			4: &vertex{0, false, 0},
			5: &vertex{0, false, 0},
			6: &vertex{0, false, 0},
			7: &vertex{0, false, 0},
			8: &vertex{0, false, 0},
		},
		[]int{0, 0, 0, 0, 1, 1, 3, 3},
	}

	got := Kosaraju(test.edges, test.edgesRev, test.vertices)
	for i := range test.want {
		if got[i] != test.want[i] {
			t.Errorf("Kosaraju(%v) == %d, want %d", test.edges, got, test.want)
		}
	}
}
