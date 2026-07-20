package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
)

type UndirectedMultiplex interface {
	Multiplex

	Layer(l int) graph.Undirected
}

func qUndirectedMultiplex(g UndirectedMultiplex, communities [][]graph.Node, weights, resolutions []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type UndirectedLayers []graph.Undirected

func NewUndirectedLayers(layers ...graph.Undirected) (UndirectedLayers, error) {
	_ = "STUB: not implemented"
	return *new(UndirectedLayers), nil
}

func (g UndirectedLayers) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g UndirectedLayers) Depth() int { _ = "STUB: not implemented"; return 0 }

func (g UndirectedLayers) Layer(l int) graph.Undirected {
	_ = "STUB: not implemented"
	return *new(graph.Undirected)
}

func louvainUndirectedMultiplex(g UndirectedMultiplex, weights, resolutions []float64, all bool, src rand.Source) *ReducedUndirectedMultiplex {
	_ = "STUB: not implemented"
	return nil
}

type ReducedUndirectedMultiplex struct {
	nodes  []multiplexCommunity
	layers []undirectedEdges

	communities [][]graph.Node

	parent *ReducedUndirectedMultiplex
}

var (
	_ UndirectedMultiplex      = (*ReducedUndirectedMultiplex)(nil)
	_ graph.WeightedUndirected = (*undirectedLayerHandle)(nil)
)

func (g *ReducedUndirectedMultiplex) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *ReducedUndirectedMultiplex) Depth() int { _ = "STUB: not implemented"; return 0 }

func (g *ReducedUndirectedMultiplex) Layer(l int) graph.Undirected {
	_ = "STUB: not implemented"
	return *new(graph.Undirected)
}

func (g *ReducedUndirectedMultiplex) Communities() [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (g *ReducedUndirectedMultiplex) Structure() [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (g *ReducedUndirectedMultiplex) Expanded() ReducedMultiplex {
	_ = "STUB: not implemented"
	return *new(ReducedMultiplex)
}

func reduceUndirectedMultiplex(g UndirectedMultiplex, communities [][]graph.Node, weights []float64) *ReducedUndirectedMultiplex {
	_ = "STUB: not implemented"
	return nil
}

type undirectedLayerHandle struct {
	multiplex *ReducedUndirectedMultiplex

	layer int
}

func (g undirectedLayerHandle) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g undirectedLayerHandle) has(id int64) bool { _ = "STUB: not implemented"; return false }

func (g undirectedLayerHandle) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g undirectedLayerHandle) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g undirectedLayerHandle) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g undirectedLayerHandle) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g undirectedLayerHandle) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g undirectedLayerHandle) EdgeBetween(xid, yid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g undirectedLayerHandle) WeightedEdgeBetween(xid, yid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g undirectedLayerHandle) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type undirectedMultiplexLocalMover struct {
	g *ReducedUndirectedMultiplex

	nodes []graph.Node

	edgeWeightOf [][]float64

	m2 []float64

	weight []func(xid, yid int64) float64

	communities [][]graph.Node

	memberships []int

	resolutions []float64

	weights []float64

	searchAll bool

	moved bool

	changed bool
}

func newUndirectedMultiplexLocalMover(g *ReducedUndirectedMultiplex, communities [][]graph.Node, weights, resolutions []float64, all bool) *undirectedMultiplexLocalMover {
	_ = "STUB: not implemented"
	return nil
}

func (l *undirectedMultiplexLocalMover) localMovingHeuristic(rnd func(int) int) (done bool) {
	_ = "STUB: not implemented"
	return false
}

func (l *undirectedMultiplexLocalMover) shuffle(rnd func(n int) int) {
	_ = "STUB: not implemented"
	return
}

func (l *undirectedMultiplexLocalMover) move(dst int, src commIdx) {
	_ = "STUB: not implemented"
	return
}

func (l *undirectedMultiplexLocalMover) deltaQ(n graph.Node) (deltaQ float64, dst int, src commIdx) {
	_ = "STUB: not implemented"
	return 0, 0, *new(commIdx)
}
