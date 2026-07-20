package iterator

import "gonum.org/v1/gonum/graph"

type OrderedEdges struct {
	idx   int
	edges []graph.Edge
}

func NewOrderedEdges(edges []graph.Edge) *OrderedEdges { _ = "STUB: not implemented"; return nil }

func (e *OrderedEdges) Len() int { _ = "STUB: not implemented"; return 0 }

func (e *OrderedEdges) Next() bool { _ = "STUB: not implemented"; return false }

func (e *OrderedEdges) Edge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (e *OrderedEdges) EdgeSlice() []graph.Edge { _ = "STUB: not implemented"; return nil }

func (e *OrderedEdges) Reset() { _ = "STUB: not implemented"; return }

type OrderedWeightedEdges struct {
	idx   int
	edges []graph.WeightedEdge
}

func NewOrderedWeightedEdges(edges []graph.WeightedEdge) *OrderedWeightedEdges {
	_ = "STUB: not implemented"
	return nil
}

func (e *OrderedWeightedEdges) Len() int { _ = "STUB: not implemented"; return 0 }

func (e *OrderedWeightedEdges) Next() bool { _ = "STUB: not implemented"; return false }

func (e *OrderedWeightedEdges) WeightedEdge() graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (e *OrderedWeightedEdges) WeightedEdgeSlice() []graph.WeightedEdge {
	_ = "STUB: not implemented"
	return nil
}

func (e *OrderedWeightedEdges) Reset() { _ = "STUB: not implemented"; return }
