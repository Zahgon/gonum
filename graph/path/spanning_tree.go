package path

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type WeightedBuilder interface {
	AddNode(graph.Node)
	SetWeightedEdge(graph.WeightedEdge)
}

func Prim(dst WeightedBuilder, g graph.WeightedUndirected) float64 {
	_ = "STUB: not implemented"
	return 0
}

type primQueue struct {
	indexOf map[int64]int
	nodes   []simple.WeightedEdge
}

func (q *primQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (q *primQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (q *primQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (q *primQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (q *primQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (q *primQueue) key(u graph.Node) (key float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (q *primQueue) update(u, v graph.Node, key float64) { _ = "STUB: not implemented"; return }

type UndirectedWeightLister interface {
	graph.WeightedUndirected
	WeightedEdges() graph.WeightedEdges
}

func Kruskal(dst WeightedBuilder, g UndirectedWeightLister) float64 {
	_ = "STUB: not implemented"
	return 0
}
