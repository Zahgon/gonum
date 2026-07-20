package path

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/linear"
	"gonum.org/v1/gonum/graph/traverse"
)

func BellmanFordFrom(u graph.Node, g traverse.Graph) (path Shortest, ok bool) {
	_ = "STUB: not implemented"
	return *new(Shortest), false
}

func BellmanFordAllFrom(u graph.Node, g traverse.Graph) (path ShortestAlts, ok bool) {
	_ = "STUB: not implemented"
	return *new(ShortestAlts), false
}

type bellmanFordQueue struct {
	queue linear.NodeQueue

	onQueue []bool

	indexOf map[int64]int
}

func (q *bellmanFordQueue) enqueue(n graph.Node) { _ = "STUB: not implemented"; return }

func (q *bellmanFordQueue) dequeue() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (q *bellmanFordQueue) len() int { _ = "STUB: not implemented"; return 0 }

func (q bellmanFordQueue) has(id int64) bool { _ = "STUB: not implemented"; return false }

func newBellmanFordQueue(indexOf map[int64]int) bellmanFordQueue {
	_ = "STUB: not implemented"
	return *new(bellmanFordQueue)
}
