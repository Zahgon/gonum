package path

import (
	"gonum.org/v1/gonum/graph"
)

func YenKShortestPaths(g graph.Graph, k int, cost float64, s, t graph.Node) [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func isSamePath(a, b []graph.Node) bool { _ = "STUB: not implemented"; return false }

type yenShortest struct {
	path   []graph.Node
	weight float64
}

type yenKSPAdjuster struct {
	graph.Graph
	isDirected bool

	weight Weighting

	visitedNodes map[int64]struct{}

	visitedEdges map[[2]int64]struct{}
}

func (g yenKSPAdjuster) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g yenKSPAdjuster) canWalk(u, v int64) bool { _ = "STUB: not implemented"; return false }

func (g yenKSPAdjuster) removeNode(u int64) { _ = "STUB: not implemented"; return }

func (g yenKSPAdjuster) removeEdge(u, v int64) { _ = "STUB: not implemented"; return }

func (g *yenKSPAdjuster) reset() { _ = "STUB: not implemented"; return }

func (g yenKSPAdjuster) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}
