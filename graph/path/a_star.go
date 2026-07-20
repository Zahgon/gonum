package path

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
)

func AStar(s, t graph.Node, g traverse.Graph, h Heuristic) (path Shortest, expanded int) {
	_ = "STUB: not implemented"
	return *new(Shortest), 0
}

func NullHeuristic(_, _ graph.Node) float64 { _ = "STUB: not implemented"; return 0 }

type aStarNode struct {
	node   graph.Node
	gscore float64
	fscore float64
}

type aStarQueue struct {
	indexOf map[int64]int
	nodes   []aStarNode
}

func (q *aStarQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (q *aStarQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (q *aStarQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (q *aStarQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (q *aStarQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (q *aStarQueue) update(id int64, g, f float64) { _ = "STUB: not implemented"; return }

func (q *aStarQueue) node(id int64) (aStarNode, bool) {
	_ = "STUB: not implemented"
	return *new(aStarNode), false
}
