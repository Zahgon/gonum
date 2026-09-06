package network

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/linear"
	"gonum.org/v1/gonum/graph/path"
)

func Betweenness(g graph.Graph) map[int64]float64 { _ = "STUB: not implemented"; return nil }

func EdgeBetweenness(g graph.Graph) map[[2]int64]float64 { _ = "STUB: not implemented"; return nil }

func brandes(g graph.Graph, accumulate func(s graph.Node, stack linear.NodeStack, p map[int64][]graph.Node, delta, sigma map[int64]float64)) {
	_ = "STUB: not implemented"
	return
}

func BetweennessWeighted(g graph.Weighted, p path.AllShortest) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

func EdgeBetweennessWeighted(g graph.Weighted, p path.AllShortest) map[[2]int64]float64 {
	_ = "STUB: not implemented"
	return nil
}
