package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
)

func qUndirected(g graph.Undirected, communities [][]graph.Node, resolution float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func louvainUndirected(g graph.Undirected, resolution float64, src rand.Source) *ReducedUndirected {
	_ = "STUB: not implemented"
	return nil
}

type ReducedUndirected struct {
	nodes []community
	undirectedEdges

	communities [][]graph.Node

	parent *ReducedUndirected
}

var (
	reducedUndirected = (*ReducedUndirected)(nil)

	_ graph.WeightedUndirected = reducedUndirected
	_ ReducedGraph             = reducedUndirected
)

func (g *ReducedUndirected) Communities() [][]graph.Node { _ = "STUB: not implemented"; return nil }

func (g *ReducedUndirected) Structure() [][]graph.Node { _ = "STUB: not implemented"; return nil }

func (g *ReducedUndirected) Expanded() ReducedGraph {
	_ = "STUB: not implemented"
	return *new(ReducedGraph)
}

func reduceUndirected(g graph.Undirected, communities [][]graph.Node) *ReducedUndirected {
	_ = "STUB: not implemented"
	return nil
}

func (g *ReducedUndirected) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *ReducedUndirected) has(id int64) bool { _ = "STUB: not implemented"; return false }

func (g *ReducedUndirected) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *ReducedUndirected) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *ReducedUndirected) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *ReducedUndirected) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *ReducedUndirected) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *ReducedUndirected) EdgeBetween(xid, yid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *ReducedUndirected) WeightedEdgeBetween(xid, yid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *ReducedUndirected) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type undirectedLocalMover struct {
	g *ReducedUndirected

	nodes []graph.Node

	edgeWeightOf []float64

	m2 float64

	weight func(xid, yid int64) float64

	communities [][]graph.Node

	memberships []int

	resolution float64

	moved bool

	changed bool
}

func newUndirectedLocalMover(g *ReducedUndirected, communities [][]graph.Node, resolution float64) *undirectedLocalMover {
	_ = "STUB: not implemented"
	return nil
}

func (l *undirectedLocalMover) localMovingHeuristic(rnd func(int) int) (done bool) {
	_ = "STUB: not implemented"
	return false
}

func (l *undirectedLocalMover) shuffle(rnd func(n int) int) { _ = "STUB: not implemented"; return }

func (l *undirectedLocalMover) move(dst int, src commIdx) { _ = "STUB: not implemented"; return }

func (l *undirectedLocalMover) deltaQ(n graph.Node) (deltaQ float64, dst int, src commIdx) {
	_ = "STUB: not implemented"
	return 0, 0, *new(commIdx)
}
