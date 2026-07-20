package topo

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

type Builder interface {
	AddNode(graph.Node)
	SetEdge(graph.Edge)
}

func CliqueGraph(dst Builder, g graph.Undirected) { _ = "STUB: not implemented"; return }

type cliqueNodeSets map[int64][]*nodeSet

type nodeSet struct {
	Clique
	nodes set.Nodes
}

type Clique struct {
	id    int64
	nodes []graph.Node
}

func (n Clique) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (n Clique) Nodes() []graph.Node { _ = "STUB: not implemented"; return nil }

type CliqueGraphEdge struct {
	from, to Clique
	nodes    []graph.Node
}

func (e CliqueGraphEdge) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e CliqueGraphEdge) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (e CliqueGraphEdge) ReversedEdge() graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (e CliqueGraphEdge) Nodes() []graph.Node { _ = "STUB: not implemented"; return nil }
