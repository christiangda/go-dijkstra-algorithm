package graph

import "testing"

func TestGraphLen(t *testing.T) {
	tests := []struct {
		name string
		g    Graph
		want int
	}{
		{name: "nil graph", g: nil, want: 0},
		{name: "empty graph", g: Graph{}, want: 0},
		{name: "nodes without edges", g: make(Graph, 3), want: 3},
		{name: "nodes with edges", g: Graph{{{To: 1, Weight: 1}}, {{To: 0, Weight: 2}}}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Len(); got != tt.want {
				t.Errorf("Len() = %d, want %d", got, tt.want)
			}
		})
	}
}
