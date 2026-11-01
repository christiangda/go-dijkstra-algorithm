package dijkstra_test

import (
	"errors"
	"fmt"

	"github.com/christiangda/go-dijkstra-algorithm/dijkstra"
	"github.com/christiangda/go-dijkstra-algorithm/graph"
)

func ExampleShortestPath() {
	//      (4.0)             (5.0)
	//    ┌────────> [1] ─────────────────┐
	//    │           ▲                   │
	//    │           │                   ▼    (2.0)
	//   [0]          │(1.0)             [3] ────────> [4]
	//    │           │                   ▲
	//    └────────> [2] ─────────────────┘
	//       (2.0)            (8.0)
	g := make(graph.Graph, 5)
	g[0] = []graph.Edge{{To: 1, Weight: 4.0}, {To: 2, Weight: 2.0}}
	g[1] = []graph.Edge{{To: 3, Weight: 5.0}}
	g[2] = []graph.Edge{{To: 1, Weight: 1.0}, {To: 3, Weight: 8.0}}
	g[3] = []graph.Edge{{To: 4, Weight: 2.0}}

	distances, previous, err := dijkstra.ShortestPath(g, 0, 4)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("cost:", distances[4])
	fmt.Println("path:", dijkstra.Path(previous, 0, 4))
	// Output:
	// cost: 10
	// path: [0 2 1 3 4]
}

// Passing a target node that does not exist disables the early exit, so the
// shortest distance to every node is computed.
func ExampleShortestPath_allNodes() {
	g := make(graph.Graph, 4)
	g[0] = []graph.Edge{{To: 1, Weight: 1.0}, {To: 2, Weight: 5.0}}
	g[1] = []graph.Edge{{To: 2, Weight: 1.0}}

	distances, _, err := dijkstra.ShortestPath(g, 0, -1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for node, distance := range distances {
		fmt.Printf("node %d: %v\n", node, distance)
	}
	// Output:
	// node 0: 0
	// node 1: 1
	// node 2: 2
	// node 3: +Inf
}

// ShortestPath rejects graphs with negative (or NaN) weights, the returned
// error wraps ErrInvalidWeight.
func ExampleShortestPath_invalidWeight() {
	g := make(graph.Graph, 2)
	g[0] = []graph.Edge{{To: 1, Weight: -3.0}}

	_, _, err := dijkstra.ShortestPath(g, 0, 1)

	fmt.Println(err)
	fmt.Println(errors.Is(err, dijkstra.ErrInvalidWeight))
	// Output:
	// dijkstra: edge weight must be non-negative: edge 0 -> 1 has weight -3
	// true
}

func ExamplePath() {
	g := make(graph.Graph, 4)
	g[0] = []graph.Edge{{To: 1, Weight: 4.0}, {To: 2, Weight: 1.0}}
	g[2] = []graph.Edge{{To: 1, Weight: 2.0}}

	_, previous, err := dijkstra.ShortestPath(g, 0, -1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dijkstra.Path(previous, 0, 1))
	fmt.Println(dijkstra.Path(previous, 0, 3)) // node 3 has no incoming edges
	// Output:
	// [0 2 1]
	// []
}
