package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

func Leiden(g graph.Graph, resolution float64, src rand.Source) ReducedGraph {
	_ = "STUB: not implemented"
	return *new(ReducedGraph)
}

const maxLeidenIterations = 1000

func leidenUndirected(g graph.Undirected, resolution float64, src rand.Source) *ReducedUndirected {
	_ = "STUB: not implemented"
	return nil
}

type inducedUndirected struct {
	g   *ReducedUndirected
	ids set.Ints[int64]
}

func (s *inducedUndirected) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (s *inducedUndirected) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedUndirected) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedUndirected) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *inducedUndirected) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (s *inducedUndirected) EdgeBetween(xid, yid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (s *inducedUndirected) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (s *inducedUndirected) WeightedEdgeBetween(xid, yid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (s *inducedUndirected) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func refineUndirected(l *undirectedLocalMover, resolution float64, rnd func(int) int) [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func LeidenMultiplex(g Multiplex, weights, resolutions []float64, all bool, src rand.Source) ReducedMultiplex {
	_ = "STUB: not implemented"
	return *new(ReducedMultiplex)
}
