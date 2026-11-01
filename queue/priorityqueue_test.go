package queue

import (
	"container/heap"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestPriorityQueuePopsInPriorityOrder(t *testing.T) {
	priorities := []float64{5, 1, 4, 1, 3, 9, 0, 2.5}

	pq := &PriorityQueue[int, float64]{}
	heap.Init(pq)

	for i, p := range priorities {
		heap.Push(pq, &Item[int, float64]{Value: i, Priority: p})
	}

	if pq.Len() != len(priorities) {
		t.Fatalf("Len() = %d, want %d", pq.Len(), len(priorities))
	}

	got := make([]float64, 0, len(priorities))
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item[int, float64])

		if item.Index != -1 {
			t.Errorf("popped item Index = %d, want -1", item.Index)
		}

		if priorities[item.Value] != item.Priority {
			t.Errorf("item %d has priority %v, want %v", item.Value, item.Priority, priorities[item.Value])
		}

		got = append(got, item.Priority)
	}

	want := slices.Clone(priorities)
	slices.Sort(want)

	if !slices.Equal(got, want) {
		t.Errorf("pop order = %v, want %v", got, want)
	}
}

func TestPriorityQueueKeepsIndexUpToDate(t *testing.T) {
	pq := &PriorityQueue[string, int]{}
	heap.Init(pq)

	for _, v := range []string{"d", "b", "e", "a", "c"} {
		heap.Push(pq, &Item[string, int]{Value: v, Priority: int(v[0])})
	}

	for i, item := range *pq {
		if item.Index != i {
			t.Errorf("item %q Index = %d, want %d", item.Value, item.Index, i)
		}
	}
}

func TestPriorityQueueInitOnExistingItems(t *testing.T) {
	pq := PriorityQueue[string, string]{
		{Value: "x", Priority: "c", Index: 0},
		{Value: "y", Priority: "a", Index: 1},
		{Value: "z", Priority: "b", Index: 2},
	}
	heap.Init(&pq)

	var got []string
	for pq.Len() > 0 {
		got = append(got, heap.Pop(&pq).(*Item[string, string]).Value)
	}

	if want := []string{"y", "z", "x"}; !slices.Equal(got, want) {
		t.Errorf("pop order = %v, want %v", got, want)
	}
}

func BenchmarkPriorityQueuePushPop(b *testing.B) {
	r := rand.New(rand.NewPCG(1, 2))

	priorities := make([]float64, 1_000)
	for i := range priorities {
		priorities[i] = r.Float64()
	}

	for b.Loop() {
		pq := &PriorityQueue[int, float64]{}

		for i, p := range priorities {
			heap.Push(pq, &Item[int, float64]{Value: i, Priority: p})
		}

		for pq.Len() > 0 {
			heap.Pop(pq)
		}
	}
}
