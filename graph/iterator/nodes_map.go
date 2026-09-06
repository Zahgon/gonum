//go:build !safe
// +build !safe

package iterator

import "gonum.org/v1/gonum/graph"

type Nodes struct {
	nodes int
	iter  *mapIter
	pos   int
	curr  graph.Node
}

func NewNodes(nodes map[int64]graph.Node) *Nodes { _ = "STUB: not implemented"; return nil }

func (n *Nodes) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *Nodes) Next() bool { _ = "STUB: not implemented"; return false }

func (n *Nodes) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (n *Nodes) Reset() { _ = "STUB: not implemented"; return }

func (n *Nodes) NodeSlice() []graph.Node { _ = "STUB: not implemented"; return nil }

type NodesByEdge struct {
	nodes map[int64]graph.Node
	edges int
	iter  *mapIter
	pos   int
	curr  graph.Node
}

func NewNodesByEdge(nodes map[int64]graph.Node, edges map[int64]graph.Edge) *NodesByEdge {
	_ = "STUB: not implemented"
	return nil
}

func NewNodesByWeightedEdge(nodes map[int64]graph.Node, edges map[int64]graph.WeightedEdge) *NodesByEdge {
	_ = "STUB: not implemented"
	return nil
}

func NewNodesByLines(nodes map[int64]graph.Node, lines map[int64]map[int64]graph.Line) *NodesByEdge {
	_ = "STUB: not implemented"
	return nil
}

func NewNodesByWeightedLines(nodes map[int64]graph.Node, lines map[int64]map[int64]graph.WeightedLine) *NodesByEdge {
	_ = "STUB: not implemented"
	return nil
}

func (n *NodesByEdge) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *NodesByEdge) Next() bool { _ = "STUB: not implemented"; return false }

func (n *NodesByEdge) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (n *NodesByEdge) Reset() { _ = "STUB: not implemented"; return }

func (n *NodesByEdge) NodeSlice() []graph.Node { _ = "STUB: not implemented"; return nil }
