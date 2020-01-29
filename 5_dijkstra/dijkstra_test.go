package main

import "testing"

func Test(t *testing.T) {
	var test = struct {
		m    map[int][]Head
		want map[int]int
	}{map[int][]Head{
		1: []Head{Head{2, 1}, Head{8, 2}},
		2: []Head{Head{1, 1}, Head{3, 1}},
		3: []Head{Head{2, 1}, Head{4, 1}},
		4: []Head{Head{3, 1}, Head{5, 1}},
		5: []Head{Head{4, 1}, Head{6, 1}},
		6: []Head{Head{5, 1}, Head{7, 1}},
		7: []Head{Head{6, 1}, Head{8, 1}},
		8: []Head{Head{7, 1}, Head{1, 2}},
	}, map[int]int{
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
		6: 4,
		7: 3,
		8: 2,
	}}
	got := Dijkstra(test.m)
	for k := range got {
		if got[k] != test.want[k] {
			t.Errorf("Dijkstra(%v) == %d, want %d", test.m, got, test.want)
		}
	}
	t.Logf("Dijkstra(%v) == %d, want %d", test.m, got, test.want)
}
