package topo

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

type Unorderable [][]graph.Node

func (e Unorderable) Error() string { _ = "STUB: not implemented"; return "" }

func lexical(nodes []graph.Node) { _ = "STUB: not implemented"; return }

func Sort(g graph.Directed) (sorted []graph.Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SortStabilized(g graph.Directed, order func([]graph.Node)) (sorted []graph.Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortedFrom(sccs [][]graph.Node, order func([]graph.Node)) ([]graph.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TarjanSCC(g graph.Directed) [][]graph.Node { _ = "STUB: not implemented"; return nil }

func tarjanSCCstabilized(g graph.Directed, order func([]graph.Node)) [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

type tarjan struct {
	succ func(id int64) []graph.Node

	index      int
	indexTable map[int64]int
	lowLink    map[int64]int
	onStack    set.Ints[int64]

	stack []graph.Node

	sccs [][]graph.Node
}

func (t *tarjan) strongconnect(v graph.Node) { _ = "STUB: not implemented"; return }
