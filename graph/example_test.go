package graph_test

import (
	"fmt"

	"github.com/christiangda/go-dijkstra-algorithm/graph"
)

func ExampleGraph() {
	//  [0] --(4.0)--> [1]
	//  [0] --(2.0)--> [2]
	//  [2] --(1.0)--> [1]
	g := make(graph.Graph, 3)
	g[0] = []graph.Edge{{To: 1, Weight: 4.0}, {To: 2, Weight: 2.0}}
	g[2] = []graph.Edge{{To: 1, Weight: 1.0}}

	fmt.Println("nodes:", g.Len())

	for from, edges := range g {
		for _, e := range edges {
			fmt.Printf("%d -> %d (%.1f)\n", from, e.To, e.Weight)
		}
	}
	// Output:
	// nodes: 3
	// 0 -> 1 (4.0)
	// 0 -> 2 (2.0)
	// 2 -> 1 (1.0)
}
