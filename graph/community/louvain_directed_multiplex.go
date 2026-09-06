package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
)

type DirectedMultiplex interface {
	Multiplex

	Layer(l int) graph.Directed
}

func qDirectedMultiplex(g DirectedMultiplex, communities [][]graph.Node, weights, resolutions []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type DirectedLayers []graph.Directed

func NewDirectedLayers(layers ...graph.Directed) (DirectedLayers, error) {
	_ = "STUB: not implemented"
	return *new(DirectedLayers), nil
}

func (g DirectedLayers) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g DirectedLayers) Depth() int { _ = "STUB: not implemented"; return 0 }

func (g DirectedLayers) Layer(l int) graph.Directed {
	_ = "STUB: not implemented"
	return *new(graph.Directed)
}

func louvainDirectedMultiplex(g DirectedMultiplex, weights, resolutions []float64, all bool, src rand.Source) *ReducedDirectedMultiplex {
	_ = "STUB: not implemented"
	return nil
}

type ReducedDirectedMultiplex struct {
	nodes  []multiplexCommunity
	layers []directedEdges

	communities [][]graph.Node

	parent *ReducedDirectedMultiplex
}

var (
	_ DirectedMultiplex      = (*ReducedDirectedMultiplex)(nil)
	_ graph.WeightedDirected = (*directedLayerHandle)(nil)
)

func (g *ReducedDirectedMultiplex) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *ReducedDirectedMultiplex) Depth() int { _ = "STUB: not implemented"; return 0 }

func (g *ReducedDirectedMultiplex) Layer(l int) graph.Directed {
	_ = "STUB: not implemented"
	return *new(graph.Directed)
}

func (g *ReducedDirectedMultiplex) Communities() [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (g *ReducedDirectedMultiplex) Structure() [][]graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (g *ReducedDirectedMultiplex) Expanded() ReducedMultiplex {
	_ = "STUB: not implemented"
	return *new(ReducedMultiplex)
}

func reduceDirectedMultiplex(g DirectedMultiplex, communities [][]graph.Node, weights []float64) *ReducedDirectedMultiplex {
	_ = "STUB: not implemented"
	return nil
}

type directedLayerHandle struct {
	multiplex *ReducedDirectedMultiplex

	layer int
}

func (g directedLayerHandle) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g directedLayerHandle) has(id int64) bool { _ = "STUB: not implemented"; return false }

func (g directedLayerHandle) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g directedLayerHandle) From(uid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g directedLayerHandle) To(vid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g directedLayerHandle) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g directedLayerHandle) HasEdgeFromTo(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g directedLayerHandle) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g directedLayerHandle) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g directedLayerHandle) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type directedMultiplexLocalMover struct {
	g *ReducedDirectedMultiplex

	nodes []graph.Node

	edgeWeightsOf [][]directedWeights

	m []float64

	weight []func(xid, yid int64) float64

	communities [][]graph.Node

	memberships []int

	resolutions []float64

	weights []float64

	searchAll bool

	moved bool

	changed bool
}

func newDirectedMultiplexLocalMover(g *ReducedDirectedMultiplex, communities [][]graph.Node, weights, resolutions []float64, all bool) *directedMultiplexLocalMover {
	_ = "STUB: not implemented"
	return nil
}

func (l *directedMultiplexLocalMover) localMovingHeuristic(rnd func(int) int) (done bool) {
	_ = "STUB: not implemented"
	return false
}

func (l *directedMultiplexLocalMover) shuffle(rnd func(n int) int) {
	_ = "STUB: not implemented"
	return
}

func (l *directedMultiplexLocalMover) move(dst int, src commIdx) { _ = "STUB: not implemented"; return }

func (l *directedMultiplexLocalMover) deltaQ(n graph.Node) (deltaQ float64, dst int, src commIdx) {
	_ = "STUB: not implemented"
	return 0, 0, *new(commIdx)
}
