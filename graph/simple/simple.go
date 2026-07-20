package simple

import (
	"gonum.org/v1/gonum/graph"
)

type Node int64

func (n Node) ID() int64 { _ = "STUB: not implemented"; return 0 }

func newSimpleNode(id int) graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

type Edge struct {
	F, T graph.Node
}

func (e Edge) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e Edge) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e Edge) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

type WeightedEdge struct {
	F, T graph.Node
	W    float64
}

func (e WeightedEdge) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e WeightedEdge) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e WeightedEdge) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (e WeightedEdge) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func isSame(a, b float64) bool { _ = "STUB: not implemented"; return false }

type edgeSetter interface {
	SetEdge(e graph.Edge)
}

type weightedEdgeSetter interface {
	SetWeightedEdge(e graph.WeightedEdge)
}
