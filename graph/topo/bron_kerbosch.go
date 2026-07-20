package topo

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

func DegeneracyOrdering(g graph.Undirected) (order []graph.Node, cores [][]graph.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KCore(k int, g graph.Undirected) []graph.Node { _ = "STUB: not implemented"; return nil }

func degeneracyOrdering(g graph.Undirected) (l []graph.Node, s []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BronKerbosch(g graph.Undirected) [][]graph.Node { _ = "STUB: not implemented"; return nil }

type bronKerbosch [][]graph.Node

func (bk *bronKerbosch) maximalCliquePivot(g graph.Undirected, r []graph.Node, p, x set.Nodes) {
	_ = "STUB: not implemented"
	return
}

func (*bronKerbosch) choosePivotFrom(g graph.Undirected, p, x set.Nodes) (neighbors []graph.Node) {
	_ = "STUB: not implemented"
	return nil
}
