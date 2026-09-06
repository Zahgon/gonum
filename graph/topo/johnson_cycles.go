package topo

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

type johnson struct {
	adjacent johnsonGraph
	b        []set.Ints[int]
	blocked  []bool
	s        int

	stack []graph.Node

	result [][]graph.Node
}

func DirectedCyclesIn(g graph.Directed) [][]graph.Node { _ = "STUB: not implemented"; return nil }

func (j *johnson) circuit(v int) bool { _ = "STUB: not implemented"; return false }

func (j *johnson) unblock(u int) { _ = "STUB: not implemented"; return }

type johnsonGraph struct {
	orig  []graph.Node
	index map[int64]int

	nodes set.Ints[int64]
	succ  map[int64]set.Ints[int64]
}

func johnsonGraphFrom(g graph.Directed) johnsonGraph {
	_ = "STUB: not implemented"
	return *new(johnsonGraph)
}

func (g johnsonGraph) order() int { _ = "STUB: not implemented"; return 0 }

func (g johnsonGraph) indexOf(id int64) int { _ = "STUB: not implemented"; return 0 }

func (g johnsonGraph) leastVertexIndex() int { _ = "STUB: not implemented"; return 0 }

func (g johnsonGraph) subgraph(s int) johnsonGraph {
	_ = "STUB: not implemented"
	return *new(johnsonGraph)
}

func (g johnsonGraph) sccSubGraph(sccs [][]graph.Node, min int) johnsonGraph {
	_ = "STUB: not implemented"
	return *new(johnsonGraph)
}

func (g johnsonGraph) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g johnsonGraph) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (johnsonGraph) Has(int64) bool { _ = "STUB: not implemented"; return false }

func (johnsonGraph) Node(int64) graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (johnsonGraph) HasEdgeBetween(_, _ int64) bool { _ = "STUB: not implemented"; return false }

func (johnsonGraph) Edge(_, _ int64) graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (johnsonGraph) HasEdgeFromTo(_, _ int64) bool { _ = "STUB: not implemented"; return false }

func (johnsonGraph) To(int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

type johnsonGraphNode int64

func (n johnsonGraphNode) ID() int64 { _ = "STUB: not implemented"; return 0 }
