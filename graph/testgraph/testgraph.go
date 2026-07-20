package testgraph

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
	"gonum.org/v1/gonum/mat"
)

func isValidIterator(it graph.Iterator) bool { _ = "STUB: not implemented"; return false }

func checkEmptyIterator(t *testing.T, it graph.Iterator, useEmpty bool) {
	_ = "STUB: not implemented"
	return
}

func hasEnds(x, y graph.Node, e Edge) bool { _ = "STUB: not implemented"; return false }

type Edge interface {
	From() graph.Node

	To() graph.Node
}

type WeightedLine interface {
	Edge

	ID() int64

	Weight() float64
}

type Builder func(nodes []graph.Node, edges []WeightedLine, self, absent float64) (g graph.Graph, n []graph.Node, e []Edge, s, a float64, ok bool)

type edgeLister interface {
	Edges() graph.Edges
}

type weightedEdgeLister interface {
	WeightedEdges() graph.WeightedEdges
}

type matrixer interface {
	Matrix() mat.Matrix
}

func ReturnAllNodes(t *testing.T, b Builder, useEmpty bool) { _ = "STUB: not implemented"; return }

func ReturnNodeSlice(t *testing.T, b Builder, useEmpty bool) { _ = "STUB: not implemented"; return }

func NodeExistence(t *testing.T, b Builder) { _ = "STUB: not implemented"; return }

func ReturnAllEdges(t *testing.T, b Builder, useEmpty bool) { _ = "STUB: not implemented"; return }

func ReturnEdgeSlice(t *testing.T, b Builder, useEmpty bool) { _ = "STUB: not implemented"; return }

func ReturnAllLines(t *testing.T, b Builder, useEmpty bool) { _ = "STUB: not implemented"; return }

func ReturnAllWeightedEdges(t *testing.T, b Builder, useEmpty bool) {
	_ = "STUB: not implemented"
	return
}

func ReturnWeightedEdgeSlice(t *testing.T, b Builder, useEmpty bool) {
	_ = "STUB: not implemented"
	return
}

func ReturnAllWeightedLines(t *testing.T, b Builder, useEmpty bool) {
	_ = "STUB: not implemented"
	return
}

func checkEdges(t *testing.T, name string, g graph.Graph, got, want []Edge) {
	_ = "STUB: not implemented"
	return
}

func EdgeExistence(t *testing.T, b Builder, reversedEdge bool) { _ = "STUB: not implemented"; return }

func LineExistence(t *testing.T, b Builder, useEmpty, reversedLine bool) {
	_ = "STUB: not implemented"
	return
}

func ReturnAdjacentNodes(t *testing.T, b Builder, useEmpty, reversedEdge bool) {
	_ = "STUB: not implemented"
	return
}

func Weight(t *testing.T, b Builder) { _ = "STUB: not implemented"; return }

func AdjacencyMatrix(t *testing.T, b Builder) { _ = "STUB: not implemented"; return }

func sortLexicalEdges(edges []Edge) { _ = "STUB: not implemented"; return }

func sortLexicalUndirectedEdges(edges []Edge) { _ = "STUB: not implemented"; return }

func cmpWeight(a, b Edge) int { _ = "STUB: not implemented"; return 0 }

func undirectedEdgeSetEqual(a, b []Edge) bool { _ = "STUB: not implemented"; return false }

func undirectedEdgeEqual(a, b Edge) bool { _ = "STUB: not implemented"; return false }

type NodeAdder interface {
	graph.Graph
	graph.NodeAdder
}

func AddNodes(t *testing.T, g NodeAdder, n int) { _ = "STUB: not implemented"; return }

func AddArbitraryNodes(t *testing.T, g NodeAdder, add graph.Nodes) {
	_ = "STUB: not implemented"
	return
}

type NodeRemover interface {
	graph.Graph
	graph.NodeRemover
}

func RemoveNodes(t *testing.T, g NodeRemover) { _ = "STUB: not implemented"; return }

type EdgeAdder interface {
	graph.Graph
	graph.EdgeAdder
}

func AddEdges(t *testing.T, n int, g EdgeAdder, newNode func(id int64) graph.Node, canLoop, canSetNode bool) {
	_ = "STUB: not implemented"
	return
}

type WeightedEdgeAdder interface {
	graph.Graph
	graph.WeightedEdgeAdder
}

func AddWeightedEdges(t *testing.T, n int, g WeightedEdgeAdder, w float64, newNode func(id int64) graph.Node, canLoop, canSetNode bool) {
	_ = "STUB: not implemented"
	return
}

func NoLoopAddEdges(t *testing.T, n int, g EdgeAdder, newNode func(id int64) graph.Node) {
	_ = "STUB: not implemented"
	return
}

func NoLoopAddWeightedEdges(t *testing.T, n int, g WeightedEdgeAdder, w float64, newNode func(id int64) graph.Node) {
	_ = "STUB: not implemented"
	return
}

type LineAdder interface {
	graph.Multigraph
	graph.LineAdder
}

func AddLines(t *testing.T, n int, g LineAdder, newNode func(id int64) graph.Node, canSetNode bool) {
	_ = "STUB: not implemented"
	return
}

type WeightedLineAdder interface {
	graph.Multigraph
	graph.WeightedLineAdder
}

func AddWeightedLines(t *testing.T, n int, g WeightedLineAdder, w float64, newNode func(id int64) graph.Node, canSetNode bool) {
	_ = "STUB: not implemented"
	return
}

type EdgeRemover interface {
	graph.Graph
	graph.EdgeRemover
}

func RemoveEdges(t *testing.T, g EdgeRemover, remove graph.Edges) {
	_ = "STUB: not implemented"
	return
}

type LineRemover interface {
	graph.Multigraph
	graph.LineRemover
}

func RemoveLines(t *testing.T, g LineRemover, remove graph.Lines) {
	_ = "STUB: not implemented"
	return
}

func undirectedIDs(e Edge) (lo, hi int64, inverted bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

type edge struct {
	f, t, id int64
}

func panics(fn func()) (ok bool) { _ = "STUB: not implemented"; return false }

type RandomNodes struct {
	n       int
	seed    uint64
	newNode func(int64) graph.Node

	curr int64

	state *rand.Rand
	seen  set.Ints[int64]
	count int
}

var _ graph.Nodes = (*RandomNodes)(nil)

func NewRandomNodes(n int, seed uint64, new func(id int64) graph.Node) *RandomNodes {
	_ = "STUB: not implemented"
	return nil
}

func (n *RandomNodes) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *RandomNodes) Next() bool { _ = "STUB: not implemented"; return false }

func (n *RandomNodes) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (n *RandomNodes) Reset() { _ = "STUB: not implemented"; return }

type tripleInt64s map[[3]int64]struct{}

func (s tripleInt64s) add(x, y, z int64) { _ = "STUB: not implemented"; return }

func (s tripleInt64s) has(x, y, z int64) bool { _ = "STUB: not implemented"; return false }
