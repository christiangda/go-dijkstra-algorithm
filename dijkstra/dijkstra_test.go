package dijkstra

import (
	"math"
	"slices"
	"testing"

	"github.com/christiangda/go-dijkstra-algorithm/graph"
)

// sampleGraph returns the same graph used by cmd/dijkstra.
func sampleGraph() graph.Graph {
	g := make(graph.Graph, 5)
	g[0] = []graph.Edge{{To: 1, Weight: 4.0}, {To: 2, Weight: 2.0}}
	g[1] = []graph.Edge{{To: 3, Weight: 5.0}}
	g[2] = []graph.Edge{{To: 1, Weight: 1.0}, {To: 3, Weight: 8.0}}
	g[3] = []graph.Edge{{To: 4, Weight: 2.0}}

	return g
}

func TestShortestPath(t *testing.T) {
	inf := math.Inf(1)

	tests := []struct {
		name          string
		graph         graph.Graph
		start, target int
		wantDistances []float64
		wantPrevious  []int
		wantPath      []int
	}{
		{
			name:          "sample graph",
			graph:         sampleGraph(),
			start:         0,
			target:        4,
			wantDistances: []float64{0, 3, 2, 8, 10},
			wantPrevious:  []int{-1, 2, 0, 1, 3},
			wantPath:      []int{0, 2, 1, 3, 4},
		},
		{
			name:          "start equals target",
			graph:         sampleGraph(),
			start:         2,
			target:        2,
			wantDistances: []float64{inf, inf, 0, inf, inf},
			wantPrevious:  []int{-1, -1, -1, -1, -1},
			wantPath:      []int{2},
		},
		{
			name:          "unreachable target",
			graph:         sampleGraph(),
			start:         4,
			target:        0,
			wantDistances: []float64{inf, inf, inf, inf, 0},
			wantPrevious:  []int{-1, -1, -1, -1, -1},
			wantPath:      nil,
		},
		{
			name:          "start out of range",
			graph:         sampleGraph(),
			start:         9,
			target:        1,
			wantDistances: []float64{inf, inf, inf, inf, inf},
			wantPrevious:  []int{-1, -1, -1, -1, -1},
			wantPath:      nil,
		},
		{
			name:          "target out of range computes every distance",
			graph:         sampleGraph(),
			start:         0,
			target:        -1,
			wantDistances: []float64{0, 3, 2, 8, 10},
			wantPrevious:  []int{-1, 2, 0, 1, 3},
			wantPath:      nil,
		},
		{
			name:          "edges to unknown nodes are ignored",
			graph:         graph.Graph{{{To: 7, Weight: 1}, {To: 1, Weight: 1}}, {}},
			start:         0,
			target:        1,
			wantDistances: []float64{0, 1},
			wantPrevious:  []int{-1, 0},
			wantPath:      []int{0, 1},
		},
		{
			name:          "empty graph",
			graph:         graph.Graph{},
			start:         0,
			target:        0,
			wantDistances: []float64{},
			wantPrevious:  []int{},
			wantPath:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distances, previous := ShortestPath(tt.graph, tt.start, tt.target)

			if !slices.Equal(distances, tt.wantDistances) {
				t.Errorf("distances = %v, want %v", distances, tt.wantDistances)
			}

			if !slices.Equal(previous, tt.wantPrevious) {
				t.Errorf("previous = %v, want %v", previous, tt.wantPrevious)
			}

			if path := Path(previous, tt.start, tt.target); !slices.Equal(path, tt.wantPath) {
				t.Errorf("Path() = %v, want %v", path, tt.wantPath)
			}
		})
	}
}

func TestPathWithCycleTerminates(t *testing.T) {
	// a malformed previous slice where 0 and 1 point at each other
	previous := []int{1, 0, -1}

	if path := Path(previous, 2, 0); path != nil {
		t.Errorf("Path() = %v, want nil", path)
	}
}
