package network

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

func MaxFlowDinic(g graph.WeightedDirected, s, t graph.Node, eps float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func initializeResidualGraph(g graph.WeightedDirected) *simple.WeightedDirectedGraph {
	_ = "STUB: not implemented"
	return nil
}

func canReachTargetInLevelGraph(r *simple.WeightedDirectedGraph, s, t graph.Node, parents [][]int64) bool {
	_ = "STUB: not implemented"
	return false
}

func computeBlockingPath(r *simple.WeightedDirectedGraph, s, t graph.Node, parents [][]int64) float64 {
	_ = "STUB: not implemented"
	return 0
}
