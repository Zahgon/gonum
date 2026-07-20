package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

func refineDirectedMultiplex(l *directedMultiplexLocalMover, weights, resolutions []float64, all bool, rnd func(int) int) [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func leidenDirectedMultiplex(g DirectedMultiplex, weights, resolutions []float64, all bool, src rand.Source) *ReducedDirectedMultiplex {
	_ = "STUB: not implemented"
	return nil
}

type inducedDirectedMultiplex struct {
	g   *ReducedDirectedMultiplex
	ids set.Ints[int64]
}

func (s *inducedDirectedMultiplex) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (s *inducedDirectedMultiplex) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedDirectedMultiplex) Depth() int { _ = "STUB: not implemented"; return 0 }

func (s *inducedDirectedMultiplex) Layer(l int) graph.Directed {
	_ = "STUB: not implemented"
	return *new(graph.Directed)
}

type inducedMultiplexLayerDirected struct {
	multiplex *ReducedDirectedMultiplex
	layer     int
	ids       set.Ints[int64]
}

func (s *inducedMultiplexLayerDirected) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (s *inducedMultiplexLayerDirected) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedMultiplexLayerDirected) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedMultiplexLayerDirected) To(vid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedMultiplexLayerDirected) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *inducedMultiplexLayerDirected) HasEdgeFromTo(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *inducedMultiplexLayerDirected) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (s *inducedMultiplexLayerDirected) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (s *inducedMultiplexLayerDirected) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}
