# 🧭 go-dijkstra-algorithm

[![Go Reference](https://pkg.go.dev/badge/github.com/christiangda/go-dijkstra-algorithm.svg)](https://pkg.go.dev/github.com/christiangda/go-dijkstra-algorithm)
[![Go Version](https://img.shields.io/github/go-mod/go-version/christiangda/go-dijkstra-algorithm)](go.mod)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

A simple, dependency-free implementation of **Dijkstra's shortest path algorithm** in Go 🐹, built on the standard library `container/heap` and Go generics.

---

## 📚 Table of contents

- [✨ Features](#-features)
- [🤔 What is Dijkstra's algorithm?](#-what-is-dijkstras-algorithm)
- [📦 Installation](#-installation)
- [🚀 Quick start](#-quick-start)
- [🖥️ Running the demo program](#️-running-the-demo-program)
- [📖 API overview](#-api-overview)
- [⚙️ How it works](#️-how-it-works)
- [🗂️ Project layout](#️-project-layout)
- [🧪 Testing and benchmarks](#-testing-and-benchmarks)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---

## ✨ Features

- ⚡ **Fast**: binary min-heap priority queue, `O((V + E) log V)`.
- 🎯 **Early exit**: stops as soon as the target node's distance is final.
- 🧹 **Lazy deletion**: no "decrease-key" needed; stale queue entries are skipped.
- 🧬 **Generic priority queue** (`queue.PriorityQueue[V, K]`) you can reuse for your own types.
- 🛡️ **Safe on bad input**: out-of-range nodes and edges never panic.
- 🪶 **Zero dependencies**: only the Go standard library.
- ✅ **100% test coverage** on the library packages, plus examples and benchmarks.

## 🤔 What is Dijkstra's algorithm?

Dijkstra's algorithm, created by **Edsger W. Dijkstra** in 1956, finds the cheapest path from a single **start node** to the other nodes of a **weighted graph**. It works as long as every edge weight is **non-negative** (`>= 0`).

> ⚠️ **Negative weights are not supported.** They are not validated and give wrong results. Use an algorithm such as Bellman-Ford for graphs with negative weights.

## 📦 Installation

Requires **Go 1.27** or newer.

```bash
go get github.com/christiangda/go-dijkstra-algorithm
```

## 🚀 Quick start

A graph is an **adjacency list**: `g[n]` holds the edges leaving node `n`. Node ids are consecutive integers starting at `0`.

```go
package main

import (
    "fmt"

    "github.com/christiangda/go-dijkstra-algorithm/dijkstra"
    "github.com/christiangda/go-dijkstra-algorithm/graph"
)

func main() {
    //  [0] --(4.0)--> [1]
    //  [0] --(1.0)--> [2] --(2.0)--> [1]
    g := make(graph.Graph, 3)
    g[0] = []graph.Edge{{To: 1, Weight: 4.0}, {To: 2, Weight: 1.0}}
    g[2] = []graph.Edge{{To: 1, Weight: 2.0}}

    distances, previous := dijkstra.ShortestPath(g, 0, 1)

    fmt.Println("cost:", distances[1])                // cost: 3
    fmt.Println("path:", dijkstra.Path(previous, 0, 1)) // path: [0 2 1]
}
```

💡 Need the distance to **every** node? Pass a target that doesn't exist, such as `-1`, to turn off the early exit:

```go
distances, previous := dijkstra.ShortestPath(g, 0, -1)
```

## 🖥️ Running the demo program

The repository includes a small command, [`cmd/godijkstra`](cmd/godijkstra/main.go), that solves this sample graph:

```text
     (4.0)             (5.0)
   ┌────────> [1] ─────────────────┐
   │           ▲                   │
   │           │                   ▼    (2.0)
  [0]          │(1.0)             [3] ────────> [4]
   │           │                   ▲
   └────────> [2] ─────────────────┘
      (2.0)            (8.0)
```

It looks for the cheapest route from node `0` to node `4`.

### ▶️ Run it directly

From the root of the repository:

```bash
go run ./cmd/godijkstra
```

### 🔨 Build a binary

```bash
go build -o godijkstra ./cmd/godijkstra
./godijkstra
```

### 🌍 Install it in your `$GOPATH/bin`

```bash
go install github.com/christiangda/go-dijkstra-algorithm/cmd/godijkstra@latest
godijkstra
```

### 📤 Expected output

```text
Cost to node 4: 10.00
Shortest path to node 4: 0->2->1->3->4
```

Going `0 → 2 → 1` (2 + 1 = 3) is cheaper than the direct edge `0 → 1` (4). Then `1 → 3 → 4` adds 5 + 2, for a total of **10**.

> 📝 The program takes no flags or arguments. To try another graph, edit the edges, `startNode` and `targetNode` in [`main.go`](cmd/godijkstra/main.go).

## 📖 API overview

Full documentation is on [pkg.go.dev](https://pkg.go.dev/github.com/christiangda/go-dijkstra-algorithm). You can also read it locally with `go doc -all ./dijkstra`.

### 🧭 `dijkstra`

| Function | Description |
| --- | --- |
| `ShortestPath(g, startNode, targetNode) (distances []float64, previous []int)` | Runs Dijkstra from `startNode`. Both slices have length `g.Len()`. `distances[n]` is the cost to reach `n` (`+Inf` if it was not reached) and `previous[n]` is the node before `n` on that path (`-1` if none). |
| `Path(previous, startNode, targetNode) []int` | Rebuilds the route `startNode → … → targetNode` from `previous`. Returns `nil` if the target is unreachable or out of range. |

> ℹ️ Because of the early exit, only `distances[targetNode]` and the distances of the nodes on its path are guaranteed to be final.

### 🕸️ `graph`

| Type | Description |
| --- | --- |
| `Edge{To int; Weight float64}` | A directed edge to node `To` with cost `Weight` (`>= 0`). |
| `Graph [][]Edge` | Adjacency list: `g[n]` are the outgoing edges of node `n`. `Len()` returns the number of nodes. |

For an **undirected** graph, add each edge in both directions.

### 📥 `queue`

A generic min-priority queue that implements `heap.Interface`. Use it through `container/heap`:

```go
pq := &queue.PriorityQueue[string, int]{}
heap.Init(pq)
heap.Push(pq, &queue.Item[string, int]{Value: "write docs", Priority: 2})
heap.Push(pq, &queue.Item[string, int]{Value: "fix bug", Priority: 1})

item := heap.Pop(pq).(*queue.Item[string, int]) // "fix bug" 🐛
```

## ⚙️ How it works

1. 🏁 Every distance starts at `+Inf` and every predecessor at `-1`. The start node's distance is set to `0` and the start node is pushed onto the priority queue.
2. 🔁 Pop the node with the smallest tentative distance.
   - 🗑️ If its recorded distance is already lower, the entry is stale: skip it (lazy deletion).
   - 🎯 If it is the target, stop (early exit).
3. 🔗 For each outgoing edge, if `distance[current] + weight` is lower than the neighbor's distance, update the neighbor's distance and predecessor, then push the neighbor onto the queue.
4. 🧵 Once the search ends, `Path` follows `previous` backwards from the target to the start and reverses the result.

| | Complexity |
| --- | --- |
| ⏱️ Time | `O((V + E) log V)` |
| 💾 Memory | `O(V + E)` |

## 🗂️ Project layout

```text
.
├── cmd/
│   └── godijkstra/      # 🖥️ demo command
│       └── main.go
├── dijkstra/            # 🧭 ShortestPath and Path
├── graph/               # 🕸️ Graph and Edge types
├── queue/               # 📥 generic priority queue
├── go.mod
├── LICENSE
└── README.md
```

## 🧪 Testing and benchmarks

Run the tests, including the [testable examples](https://go.dev/blog/examples):

```bash
go test ./...
```

With the race detector and coverage:

```bash
go test -race -cover ./...
```

Run the benchmarks. They use reproducible random graphs of 100 to 100,000 nodes:

```bash
go test -run '^$' -bench . -benchmem ./...
```

## 🤝 Contributing

Contributions are welcome! 🎉

1. 🍴 Fork the repository.
2. 🌱 Create a branch: `git checkout -b feat/my-feature`.
3. ✅ Make sure `gofmt -l .`, `go vet ./...` and `go test ./...` pass.
4. 📬 Open a pull request.

## 📄 License

Distributed under the **Apache License 2.0**. See [LICENSE](LICENSE) for details.

---

Made with ❤️ and 🐹 Go.
