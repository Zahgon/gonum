package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
)

func qDirected(g graph.Directed, communities [][]graph.Node, resolution float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func louvainDirected(g graph.Directed, resolution float64, src rand.Source) ReducedGraph {
	_ = "STUB: not implemented"
	return *new(ReducedGraph)
}

type ReducedDirected struct {
	nodes []community
	directedEdges

	communities [][]graph.Node

	parent *ReducedDirected
}

var (
	reducedDirected = (*ReducedDirected)(nil)

	_ graph.WeightedDirected = reducedDirected
	_ ReducedGraph           = reducedDirected
)

func (g *ReducedDirected) Communities() [][]graph.Node { _ = "STUB: not implemented"; return nil }

func (g *ReducedDirected) Structure() [][]graph.Node { _ = "STUB: not implemented"; return nil }

func (g *ReducedDirected) Expanded() ReducedGraph {
	_ = "STUB: not implemented"
	return *new(ReducedGraph)
}

func reduceDirected(g graph.Directed, communities [][]graph.Node) *ReducedDirected {
	_ = "STUB: not implemented"
	return nil
}

func (g *ReducedDirected) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *ReducedDirected) has(id int64) bool { _ = "STUB: not implemented"; return false }

func (g *ReducedDirected) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *ReducedDirected) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *ReducedDirected) To(vid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *ReducedDirected) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *ReducedDirected) HasEdgeFromTo(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *ReducedDirected) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *ReducedDirected) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *ReducedDirected) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type directedLocalMover struct {
	g *ReducedDirected

	nodes []graph.Node

	edgeWeightsOf []directedWeights

	m float64

	weight func(xid, yid int64) float64

	communities [][]graph.Node

	memberships []int

	resolution float64

	moved bool

	changed bool
}

type directedWeights struct {
	out, in float64
}

func newDirectedLocalMover(g *ReducedDirected, communities [][]graph.Node, resolution float64) *directedLocalMover {
	_ = "STUB: not implemented"
	return nil
}

func (l *directedLocalMover) localMovingHeuristic(rnd func(int) int) (done bool) {
	_ = "STUB: not implemented"
	return false
}

func (l *directedLocalMover) shuffle(rnd func(n int) int) { _ = "STUB: not implemented"; return }

func (l *directedLocalMover) move(dst int, src commIdx) { _ = "STUB: not implemented"; return }

func (l *directedLocalMover) deltaQ(n graph.Node) (deltaQ float64, dst int, src commIdx) {
	_ = "STUB: not implemented"
	return 0, 0, *new(commIdx)
}
