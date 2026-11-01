// Command dijkstra is a small demo of the dijkstra package.
//
// It builds a sample directed, weighted graph of five nodes, finds the
// cheapest route from node 0 to node 4 and prints its cost and path:
//
//	$ go run ./cmd/dijkstra
//	Cost to node 4: 10.00
//	Shortest path to node 4: 0->2->1->3->4
//
// The program takes no flags or arguments; edit the graph in main to try
// other inputs.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/christiangda/go-dijkstra-algorithm/dijkstra"
	"github.com/christiangda/go-dijkstra-algorithm/graph"
)

func main() {
	//      (4.0)             (5.0)
	//    ┌────────> [1] ─────────────────┐
	//    │           ▲                   │
	//    │           │                   ▼    (2.0)
	//   [0]          │(1.0)             [3] ────────> [4]
	//    │           │                   ▲
	//    └────────> [2] ─────────────────┘
	//       (2.0)            (8.0)

	g := make(graph.Graph, 5)
	g[0] = append(g[0], graph.Edge{To: 1, Weight: 4.0}, graph.Edge{To: 2, Weight: 2.0})
	g[1] = append(g[1], graph.Edge{To: 3, Weight: 5.0})
	g[2] = append(g[2], graph.Edge{To: 1, Weight: 1.0}, graph.Edge{To: 3, Weight: 8.0})
	g[3] = append(g[3], graph.Edge{To: 4, Weight: 2.0})

	startNode := 0
	targetNode := 4

	distances, previous, err := dijkstra.ShortestPath(g, startNode, targetNode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	path := dijkstra.Path(previous, startNode, targetNode)

	asArrows := func(path []int) string {
		var sb strings.Builder

		for i, p := range path {
			sb.WriteString(strconv.Itoa(p))

			if i < len(path)-1 {
				sb.WriteString("->")
			}
		}

		return sb.String()
	}

	fmt.Printf("Cost to node %d: %.2f\n", targetNode, distances[targetNode])
	fmt.Printf("Shortest path to node %v: %v\n", targetNode, asArrows(path))

	// Cost to node 4: 10.00
	// Shortest path to node 4: 0->2->1->3->4
}
