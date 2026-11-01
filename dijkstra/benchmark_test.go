package dijkstra

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/christiangda/go-dijkstra-algorithm/graph"
)

// randomGraph builds a reproducible graph of n nodes where every node has
// edgesPerNode outgoing edges to random nodes with weights in [0, 100).
// A chain 0 -> 1 -> ... -> n-1 is always included so the last node is
// reachable from node 0.
func randomGraph(n, edgesPerNode int) graph.Graph {
	r := rand.New(rand.NewPCG(42, 1024))

	g := make(graph.Graph, n)
	for from := range n {
		g[from] = make([]graph.Edge, 0, edgesPerNode+1)

		if from < n-1 {
			g[from] = append(g[from], graph.Edge{To: from + 1, Weight: 100})
		}

		for range edgesPerNode {
			g[from] = append(g[from], graph.Edge{To: r.IntN(n), Weight: r.Float64() * 100})
		}
	}

	return g
}

func BenchmarkShortestPath(b *testing.B) {
	for _, n := range []int{100, 1_000, 10_000, 100_000} {
		g := randomGraph(n, 5)

		b.Run(fmt.Sprintf("nodes=%d/target", n), func(b *testing.B) {
			for b.Loop() {
				ShortestPath(g, 0, n-1)
			}
		})

		b.Run(fmt.Sprintf("nodes=%d/all", n), func(b *testing.B) {
			for b.Loop() {
				ShortestPath(g, 0, -1)
			}
		})
	}
}

func BenchmarkPath(b *testing.B) {
	for _, n := range []int{1_000, 100_000} {
		g := randomGraph(n, 5)
		_, previous, err := ShortestPath(g, 0, -1)
		if err != nil {
			b.Fatal(err)
		}

		b.Run(fmt.Sprintf("nodes=%d", n), func(b *testing.B) {
			for b.Loop() {
				Path(previous, 0, n-1)
			}
		})
	}
}
