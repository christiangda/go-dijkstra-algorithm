package queue_test

import (
	"container/heap"
	"fmt"

	"github.com/christiangda/go-dijkstra-algorithm/queue"
)

func ExamplePriorityQueue() {
	pq := &queue.PriorityQueue[string, int]{}
	heap.Init(pq)

	heap.Push(pq, &queue.Item[string, int]{Value: "write docs", Priority: 3})
	heap.Push(pq, &queue.Item[string, int]{Value: "fix bug", Priority: 1})
	heap.Push(pq, &queue.Item[string, int]{Value: "add tests", Priority: 2})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*queue.Item[string, int])
		fmt.Printf("%d: %s\n", item.Priority, item.Value)
	}
	// Output:
	// 1: fix bug
	// 2: add tests
	// 3: write docs
}
