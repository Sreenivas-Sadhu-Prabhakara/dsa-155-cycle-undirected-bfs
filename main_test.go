package main

import "testing"

func TestHasCycle(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  bool
	}{
		{3, [][2]int{{0, 1}, {1, 2}, {2, 0}}, true},         // triangle
		{3, [][2]int{{0, 1}, {1, 2}}, false},                // path
		{4, [][2]int{{0, 1}, {2, 3}}, false},                // two components, no cycle
		{5, [][2]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}, true}, // cycle in the second component
		{1, nil, false},
		{0, nil, false},
		{2, [][2]int{{0, 1}}, false}, // single edge is not a cycle
	}
	for _, c := range cases {
		if got := HasCycle(c.n, c.edges); got != c.want {
			t.Errorf("HasCycle(n=%d, %v) = %v, want %v", c.n, c.edges, got, c.want)
		}
	}
}
