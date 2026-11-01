// Package queue implements a generic min-priority queue on top of the
// standard library [container/heap] package.
//
// [PriorityQueue] satisfies [heap.Interface], so it must be used through the
// functions of container/heap, never by calling its methods directly:
//
//	pq := &queue.PriorityQueue[string, int]{}
//	heap.Init(pq)
//	heap.Push(pq, &queue.Item[string, int]{Value: "b", Priority: 2})
//	heap.Push(pq, &queue.Item[string, int]{Value: "a", Priority: 1})
//
//	item := heap.Pop(pq).(*queue.Item[string, int]) // item.Value == "a"
//
// The item with the lowest Priority is always popped first.
package queue

import "container/heap"

// Compile-time check that *PriorityQueue satisfies heap.Interface.
var _ heap.Interface = (*PriorityQueue[int, float64])(nil)

// Item is an element of a [PriorityQueue].
//
// V is the type of the stored value and K the type of its priority.
type Item[V, K comparable] struct {
	// Value is the payload stored in the queue, for example a node id.
	Value V

	// Priority decides the order of the item, lower values are popped first.
	Priority K

	// Index is the position of the item in the heap. It is maintained by the
	// heap.Interface methods and is set to -1 once the item is popped.
	Index int
}

// PriorityQueue is a min-heap of *[Item] ordered by Priority.
// It implements [heap.Interface].
type PriorityQueue[V comparable, K Ordered] []*Item[V, K]

// Len returns the number of items in the priority queue.
// It is part of [sort.Interface].
func (pq PriorityQueue[V, K]) Len() int {
	return len(pq)
}

// Less reports whether the item at index i has a lower priority than the item
// at index j, which makes the queue a min-heap.
// It is part of [sort.Interface].
func (pq PriorityQueue[V, K]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the items at indexes i and j and updates their Index.
// It is part of [sort.Interface].
func (pq PriorityQueue[V, K]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push appends x, which must be an *Item[V, K], to the end of the queue.
// It is part of [heap.Interface]; call [heap.Push] instead of using it directly.
func (pq *PriorityQueue[V, K]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[V, K])
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the last item of the queue as an *Item[V, K].
// It is part of [heap.Interface]; call [heap.Pop] instead, which returns the
// item with the lowest priority.
func (pq *PriorityQueue[V, K]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.Index = -1 // for safety
	*pq = old[0 : n-1]

	return item
}
