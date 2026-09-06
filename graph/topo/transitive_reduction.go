package topo

import (
	"gonum.org/v1/gonum/graph"
)

type DirectedGraphRemover interface {
	graph.Directed
	graph.EdgeRemover
}

func TransitiveReduce(g DirectedGraphRemover) { _ = "STUB: not implemented"; return }

func idsFrom(it graph.Nodes) []int64 { _ = "STUB: not implemented"; return nil }

func indexNodes(g graph.Graph) (ids []int64, indexOf map[int64]int) {
	_ = "STUB: not implemented"
	return nil, nil
}
