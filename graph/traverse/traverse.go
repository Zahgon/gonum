package traverse

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/linear"
	"gonum.org/v1/gonum/graph/internal/set"
)

var _ Graph = graph.Graph(nil)

type Graph interface {
	From(id int64) graph.Nodes

	Edge(uid, vid int64) graph.Edge
}

type BreadthFirst struct {
	Visit func(graph.Node)

	Traverse func(graph.Edge) bool

	queue   linear.NodeQueue
	visited set.Ints[int64]
}

func (b *BreadthFirst) Walk(g Graph, from graph.Node, until func(n graph.Node, d int) bool) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (b *BreadthFirst) WalkAll(g graph.Undirected, before, after func(), during func(graph.Node)) {
	_ = "STUB: not implemented"
	return
}

func (b *BreadthFirst) Visited(n graph.Node) bool { _ = "STUB: not implemented"; return false }

func (b *BreadthFirst) Reset() { _ = "STUB: not implemented"; return }

type DepthFirst struct {
	Visit func(graph.Node)

	Traverse func(graph.Edge) bool

	stack   linear.NodeStack
	visited set.Ints[int64]
}

func (d *DepthFirst) Walk(g Graph, from graph.Node, until func(graph.Node) bool) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (d *DepthFirst) WalkAll(g graph.Undirected, before, after func(), during func(graph.Node)) {
	_ = "STUB: not implemented"
	return
}

func (d *DepthFirst) Visited(n graph.Node) bool { _ = "STUB: not implemented"; return false }

func (d *DepthFirst) Reset() { _ = "STUB: not implemented"; return }
