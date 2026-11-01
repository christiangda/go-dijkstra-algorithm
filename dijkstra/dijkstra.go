// Package dijkstra implements Dijkstra's shortest path algorithm.
//
// Dijkstra's algorithm (created by Edsger W. Dijkstra in 1956) solves the
// single-source shortest path problem: given a start node, it finds the
// cheapest path to every other node of a weighted graph whose edge weights
// are all non-negative (>= 0).
//
// # Usage
//
// Build a [graph.Graph], call [ShortestPath] to compute the distances and the
// predecessor of each node, then call [Path] to rebuild the route to the node
// you are interested in:
//
//	g := make(graph.Graph, 3)
//	g[0] = []graph.Edge{{To: 1, Weight: 4}, {To: 2, Weight: 1}}
//	g[2] = []graph.Edge{{To: 1, Weight: 2}}
//
//	distances, previous := dijkstra.ShortestPath(g, 0, 1)
//	path := dijkstra.Path(previous, 0, 1) // [0 2 1], cost distances[1] == 3
//
// # Implementation
//
// The frontier of nodes to visit is kept in a binary min-heap
// ([queue.PriorityQueue]) ordered by the tentative distance, which gives a
// running time of O((V + E) log V) for V nodes and E edges.
//
// Instead of a "decrease-key" operation, a node is pushed again every time a
// cheaper distance to it is found, and stale entries are skipped when they are
// popped (lazy deletion). The search also stops as soon as the target node is
// popped from the heap (early exit), because at that moment its distance is
// final.
package dijkstra

import (
	"container/heap"
	"math"
	"slices"

	"github.com/christiangda/go-dijkstra-algorithm/graph"
	"github.com/christiangda/go-dijkstra-algorithm/queue"
)

// ShortestPath runs Dijkstra's algorithm on g from startNode and stops as soon
// as the shortest distance to targetNode is known.
//
// Both returned slices have length g.Len() and are indexed by node id:
//
//   - distances[n] is the cost of the cheapest path from startNode to n, or
//     +Inf ([math.Inf](1)) when n was not reached.
//   - previous[n] is the node that comes right before n on that path, or -1
//     when n is startNode or was not reached. Pass it to [Path] to rebuild
//     the full route.
//
// Because of the early exit, only distances[targetNode] (and the distances of
// the nodes along its path) are guaranteed to be final; other nodes may hold
// a value that is higher than their real shortest distance. To compute the
// shortest distance to every node, pass a targetNode that does not exist in
// the graph, for example -1.
//
// If startNode is not a valid node id, every distance is +Inf and every
// previous entry is -1. Edges that point to a node outside of the graph are
// ignored.
//
// Edge weights must be non-negative. Negative weights are not validated and
// produce incorrect results; use an algorithm such as Bellman-Ford for them.
func ShortestPath(g graph.Graph, startNode, targetNode int) (distances []float64, previous []int) {
	graphLength := g.Len()
	distances = make([]float64, graphLength)
	previous = make([]int, graphLength)

	// initialize distances with infinite (math.Inf) and previous with -1
	for i := range graphLength {
		distances[i] = math.Inf(1)
		previous[i] = -1
	}

	if startNode < 0 || startNode >= graphLength {
		return distances, previous
	}

	distances[startNode] = 0.0

	// the queue holds node ids (Value) ordered by their tentative distance (Priority)
	pq := &queue.PriorityQueue[int, float64]{}
	heap.Init(pq)
	heap.Push(pq, &queue.Item[int, float64]{Value: startNode, Priority: 0.0})

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*queue.Item[int, float64])
		currentNode := currentItem.Value
		currentDist := currentItem.Priority

		// Lazy Deletion: if we found an obsolete item in the heap, we ignore it
		if currentDist > distances[currentNode] {
			continue
		}

		// Early Exit: if we arrive to target, we finished
		if currentNode == targetNode {
			break
		}

		for _, edge := range g[currentNode] {
			nextNode := edge.To
			weight := edge.Weight

			if nextNode < 0 || nextNode >= graphLength {
				continue
			}

			newDistance := distances[currentNode] + weight

			if newDistance < distances[nextNode] {
				distances[nextNode] = newDistance
				previous[nextNode] = currentNode
				heap.Push(pq, &queue.Item[int, float64]{Value: nextNode, Priority: newDistance})
			}
		}
	}

	return distances, previous
}

// Path rebuilds the shortest route from startNode to targetNode using the
// previous slice returned by [ShortestPath].
//
// The returned slice lists the node ids in travel order, starting with
// startNode and ending with targetNode. When startNode equals targetNode the
// path is just []int{startNode}.
//
// previous must come from a call to [ShortestPath] with the same startNode.
// Path returns nil when targetNode is out of range or is not reachable from
// startNode.
func Path(previous []int, startNode, targetNode int) (path []int) {
	if targetNode < 0 || targetNode >= len(previous) {
		return nil
	}

	// if the targetNode if not reachable and it is not startNode
	if previous[targetNode] == -1 && targetNode != startNode {
		return nil
	}

	path = make([]int, 0, len(previous))
	curr := targetNode

	// we traverse in reverse using the`previous` slice, a valid path never has
	// more than len(previous) nodes, so the bound protects against cycles
	for curr != -1 && len(path) < len(previous) {
		path = append(path, curr)
		if curr == startNode {
			break
		}

		curr = previous[curr]
	}

	// are we in the startNode?
	if len(path) == 0 || path[len(path)-1] != startNode {
		return nil // there is not valid path
	}

	slices.Reverse(path)

	return path
}
