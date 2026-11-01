// Package graph has the data structures necessary to define a graph
package graph

// Edge represent each edge in the Graph
type Edge struct {
	// To is the destination node id
	To int

	// Weight always should be greater or equal to zero, never negative
	// since Dijkstra's Algorithm does not support negative distances
	Weight float64
}

// Graph is represented for each Node id and the slice of adjacent edges
type Graph [][]Edge

func (g Graph) Len() int {
	return len(g)
}
