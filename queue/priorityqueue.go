// Package queue has the necessary data structures and method to implement a priority queue using the
// go generics
package queue

// Item represent the an element of the priority queue.
// Implement generics type to support different data type during the implementation
type Item[V, K comparable] struct {
	Value    V
	Priority K
	Index    int
}

// PriorityQueue implement heap.Interface and contains a slice *Item.
type PriorityQueue[V comparable, K Ordered] []*Item[V, K]

// Len of the priority queue, satisfies sort.Interface
func (pq PriorityQueue[V, K]) Len() int {
	return len(pq)
}

// Less reports whether the element with index i must sort before the element with index j of the priority queue,
// satisfies sort.Interface
func (pq PriorityQueue[V, K]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the elements with indexes i and j of the priority queue
// satisfies sort.Interface
func (pq PriorityQueue[V, K]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push add a Item on a priority queue
// satisfies heap.Interface
func (pq *PriorityQueue[V, K]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[V, K])
	item.Index = n
	*pq = append(*pq, item)
}

// Pop the first Item of the priority queue
// satisfies heap.Interface
func (pq *PriorityQueue[V, K]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.Index = -1 // for safety
	*pq = old[0 : n-1]

	return item
}
