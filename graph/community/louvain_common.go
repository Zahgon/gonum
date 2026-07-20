package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

func Q(g graph.Graph, communities [][]graph.Node, resolution float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

type ReducedGraph interface {
	graph.Graph

	Communities() [][]graph.Node

	Structure() [][]graph.Node

	Expanded() ReducedGraph
}

func Modularize(g graph.Graph, resolution float64, src rand.Source) ReducedGraph {
	_ = "STUB: not implemented"
	return *new(ReducedGraph)
}

type Multiplex interface {
	Nodes() graph.Nodes

	Depth() int
}

func QMultiplex(g Multiplex, communities [][]graph.Node, weights, resolutions []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type ReducedMultiplex interface {
	Multiplex

	Communities() [][]graph.Node

	Structure() [][]graph.Node

	Expanded() ReducedMultiplex
}

func ModularizeMultiplex(g Multiplex, weights, resolutions []float64, all bool, src rand.Source) ReducedMultiplex {
	_ = "STUB: not implemented"
	return *new(ReducedMultiplex)
}

type undirectedEdges struct {
	edges   [][]int
	weights map[[2]int]float64
}

type directedEdges struct {
	edgesFrom [][]int
	edgesTo   [][]int
	weights   map[[2]int]float64
}

func isValidID(id int64) bool { _ = "STUB: not implemented"; return false }

type community struct {
	id int

	nodes []graph.Node

	weight float64
}

func (n community) ID() int64 { _ = "STUB: not implemented"; return 0 }

type edge struct {
	from, to community
	weight   float64
}

func (e edge) From() graph.Node         { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e edge) To() graph.Node           { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e edge) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }
func (e edge) Weight() float64          { _ = "STUB: not implemented"; return 0 }

type multiplexCommunity struct {
	id int

	nodes []graph.Node

	weights []float64
}

func (n multiplexCommunity) ID() int64 { _ = "STUB: not implemented"; return 0 }

type multiplexEdge struct {
	from, to multiplexCommunity
	weight   float64
}

func (e multiplexEdge) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e multiplexEdge) To() graph.Node   { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e multiplexEdge) ReversedEdge() graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}
func (e multiplexEdge) Weight() float64 { _ = "STUB: not implemented"; return 0 }

type commIdx struct {
	community int
	node      int
}

type node int

func (n node) ID() int64 { _ = "STUB: not implemented"; return 0 }

type minTaker interface {
	TakeMin(p *int) bool
}

type dense struct {
	pos int
	n   int
}

func (d *dense) TakeMin(p *int) bool { _ = "STUB: not implemented"; return false }

type slice struct {
	pos   int
	elems []int
}

func newSlice(s set.Ints[int]) *slice { _ = "STUB: not implemented"; return nil }

func (s *slice) TakeMin(p *int) bool { _ = "STUB: not implemented"; return false }

const (
	negativeWeight = "community: unexpected negative edge weight"
	positiveWeight = "community: unexpected positive edge weight"

	deltaQtol = 1e-14
)

func positiveWeightFuncFor(g graph.Graph) func(xid, yid int64) float64 {
	_ = "STUB: not implemented"
	return nil
}

func negativeWeightFuncFor(g graph.Graph) func(xid, yid int64) float64 {
	_ = "STUB: not implemented"
	return nil
}

func depth(weights []float64) int { _ = "STUB: not implemented"; return 0 }
