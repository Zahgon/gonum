package multi

import "gonum.org/v1/gonum/graph"

type Node int64

func (n Node) ID() int64 { _ = "STUB: not implemented"; return 0 }

type Edge struct {
	F, T graph.Node

	graph.Lines
}

func (e Edge) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e Edge) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e Edge) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

type Line struct {
	F, T graph.Node

	UID int64
}

func (l Line) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (l Line) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (l Line) ReversedLine() graph.Line { _ = "STUB: not implemented"; return *new(graph.Line) }

func (l Line) ID() int64 { _ = "STUB: not implemented"; return 0 }

type WeightedEdge struct {
	F, T graph.Node

	graph.WeightedLines

	WeightFunc func(graph.WeightedLines) float64
}

func (e WeightedEdge) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e WeightedEdge) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e WeightedEdge) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (e WeightedEdge) Weight() float64 { _ = "STUB: not implemented"; return 0 }

type WeightedLine struct {
	F, T graph.Node
	W    float64

	UID int64
}

func (l WeightedLine) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (l WeightedLine) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (l WeightedLine) ReversedLine() graph.Line { _ = "STUB: not implemented"; return *new(graph.Line) }

func (l WeightedLine) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (l WeightedLine) Weight() float64 { _ = "STUB: not implemented"; return 0 }
