package path

import (
	"gonum.org/v1/gonum/graph"
)

func JohnsonAllPaths(g graph.Graph) (paths AllShortest, ok bool) {
	_ = "STUB: not implemented"
	return *new(AllShortest), false
}

type johnsonWeightAdjuster struct {
	graph.Graph
	weight Weighting

	adjustBy Shortest
}

var _ graph.Weighted = johnsonWeightAdjuster{}

func (g johnsonWeightAdjuster) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g johnsonWeightAdjuster) WeightedEdge(_, _ int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g johnsonWeightAdjuster) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (johnsonWeightAdjuster) HasEdgeBetween(_, _ int64) bool {
	_ = "STUB: not implemented"
	return false
}

type johnsonReWeight struct {
	johnsonWeightAdjuster
	q int64
}

func (g johnsonReWeight) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g johnsonReWeight) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g johnsonReWeight) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g johnsonReWeight) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g johnsonReWeight) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type johnsonGraphNode int64

func (n johnsonGraphNode) ID() int64 { _ = "STUB: not implemented"; return 0 }

func newJohnsonNodeIterator(q int64, nodes graph.Nodes) *johnsonNodeIterator {
	_ = "STUB: not implemented"
	return nil
}

type johnsonNodeIterator struct {
	q          int64
	nodes      graph.Nodes
	qUsed, qOK bool
}

func (it *johnsonNodeIterator) Len() int { _ = "STUB: not implemented"; return 0 }

func (it *johnsonNodeIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (it *johnsonNodeIterator) Node() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (it *johnsonNodeIterator) Reset() { _ = "STUB: not implemented"; return }
