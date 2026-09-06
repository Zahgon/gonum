package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

func leidenUndirectedMultiplex(g UndirectedMultiplex, weights, resolutions []float64, all bool, src rand.Source) *ReducedUndirectedMultiplex {
	_ = "STUB: not implemented"
	return nil
}

type inducedUndirectedMultiplex struct {
	g   *ReducedUndirectedMultiplex
	ids set.Ints[int64]
}

func (s *inducedUndirectedMultiplex) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (s *inducedUndirectedMultiplex) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedUndirectedMultiplex) Depth() int { _ = "STUB: not implemented"; return 0 }

func (s *inducedUndirectedMultiplex) Layer(l int) graph.Undirected {
	_ = "STUB: not implemented"
	return *new(graph.Undirected)
}

type inducedMultiplexLayer struct {
	multiplex *ReducedUndirectedMultiplex
	layer     int
	ids       set.Ints[int64]
}

func (s *inducedMultiplexLayer) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (s *inducedMultiplexLayer) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedMultiplexLayer) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedMultiplexLayer) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *inducedMultiplexLayer) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (s *inducedMultiplexLayer) EdgeBetween(xid, yid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (s *inducedMultiplexLayer) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (s *inducedMultiplexLayer) WeightedEdgeBetween(xid, yid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (s *inducedMultiplexLayer) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func refineUndirectedMultiplex(l *undirectedMultiplexLocalMover, weights, resolutions []float64, all bool, rnd func(int) int) [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}
