// Package graph defines the data structures used to describe a directed,
// weighted graph as an adjacency list.
//
// Nodes are identified by consecutive integers starting at 0, and a
// [Graph] stores, for each node id, the list of edges leaving that node:
//
//	//  [0] --(4.0)--> [1]
//	//  [0] --(2.0)--> [2]
//	//  [2] --(1.0)--> [1]
//	g := make(graph.Graph, 3)
//	g[0] = []graph.Edge{{To: 1, Weight: 4.0}, {To: 2, Weight: 2.0}}
//	g[2] = []graph.Edge{{To: 1, Weight: 1.0}}
//
// Nodes without outgoing edges, like node 1 above, simply keep a nil slice.
// For an undirected graph add each edge in both directions.
package graph

// Edge is a directed, weighted connection to another node of a [Graph].
// The source node is implied by the position of the edge in the Graph.
type Edge struct {
	// To is the id of the destination node.
	To int

	// Weight is the cost of traveling through the edge. It must be greater
	// than or equal to zero, because Dijkstra's algorithm does not support
	// negative weights; dijkstra.ShortestPath returns an error otherwise.
	Weight float64
}

// Graph is a directed, weighted graph represented as an adjacency list.
//
// The index of the outer slice is the node id and the inner slice holds the
// edges leaving that node, so g[n] are the outgoing edges of node n.
type Graph [][]Edge

// Len returns the number of nodes in the graph.
func (g Graph) Len() int {
	return len(g)
}
