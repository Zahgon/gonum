package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

func leidenDirected(g graph.Directed, resolution float64, src rand.Source) *ReducedDirected {
	_ = "STUB: not implemented"
	return nil
}

type inducedDirected struct {
	g   *ReducedDirected
	ids set.Ints[int64]
}

func (s *inducedDirected) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (s *inducedDirected) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (s *inducedDirected) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedDirected) To(vid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (s *inducedDirected) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *inducedDirected) HasEdgeFromTo(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *inducedDirected) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (s *inducedDirected) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (s *inducedDirected) WeightedEdgeBetween(xid, yid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (s *inducedDirected) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func refineDirected(l *directedLocalMover, resolution float64, rnd func(int) int) [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}
