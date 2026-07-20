package iterator

import "gonum.org/v1/gonum/graph"

type OrderedNodes struct {
	idx   int
	nodes []graph.Node
}

func NewOrderedNodes(nodes []graph.Node) *OrderedNodes { _ = "STUB: not implemented"; return nil }

func (n *OrderedNodes) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *OrderedNodes) Next() bool { _ = "STUB: not implemented"; return false }

func (n *OrderedNodes) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (n *OrderedNodes) NodeSlice() []graph.Node { _ = "STUB: not implemented"; return nil }

func (n *OrderedNodes) Reset() { _ = "STUB: not implemented"; return }

type LazyOrderedNodes struct {
	iter  OrderedNodes
	nodes map[int64]graph.Node
}

func NewLazyOrderedNodes(nodes map[int64]graph.Node) *LazyOrderedNodes {
	_ = "STUB: not implemented"
	return nil
}

func (n *LazyOrderedNodes) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *LazyOrderedNodes) Next() bool { _ = "STUB: not implemented"; return false }

func (n *LazyOrderedNodes) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (n *LazyOrderedNodes) NodeSlice() []graph.Node { _ = "STUB: not implemented"; return nil }

func (n *LazyOrderedNodes) Reset() { _ = "STUB: not implemented"; return }

func (n *LazyOrderedNodes) fillSlice() { _ = "STUB: not implemented"; return }

type LazyOrderedNodesByEdge struct {
	iter  OrderedNodes
	nodes map[int64]graph.Node
	edges map[int64]graph.Edge
}

func NewLazyOrderedNodesByEdge(nodes map[int64]graph.Node, edges map[int64]graph.Edge) *LazyOrderedNodesByEdge {
	_ = "STUB: not implemented"
	return nil
}

func (n *LazyOrderedNodesByEdge) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *LazyOrderedNodesByEdge) Next() bool { _ = "STUB: not implemented"; return false }

func (n *LazyOrderedNodesByEdge) Node() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (n *LazyOrderedNodesByEdge) NodeSlice() []graph.Node { _ = "STUB: not implemented"; return nil }

func (n *LazyOrderedNodesByEdge) Reset() { _ = "STUB: not implemented"; return }

func (n *LazyOrderedNodesByEdge) fillSlice() { _ = "STUB: not implemented"; return }

type LazyOrderedNodesByWeightedEdge struct {
	iter  OrderedNodes
	nodes map[int64]graph.Node
	edges map[int64]graph.WeightedEdge
}

func NewLazyOrderedNodesByWeightedEdge(nodes map[int64]graph.Node, edges map[int64]graph.WeightedEdge) *LazyOrderedNodesByWeightedEdge {
	_ = "STUB: not implemented"
	return nil
}

func (n *LazyOrderedNodesByWeightedEdge) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *LazyOrderedNodesByWeightedEdge) Next() bool { _ = "STUB: not implemented"; return false }

func (n *LazyOrderedNodesByWeightedEdge) Node() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (n *LazyOrderedNodesByWeightedEdge) NodeSlice() []graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (n *LazyOrderedNodesByWeightedEdge) Reset() { _ = "STUB: not implemented"; return }

func (n *LazyOrderedNodesByWeightedEdge) fillSlice() { _ = "STUB: not implemented"; return }

type LazyOrderedNodesByLines struct {
	iter  OrderedNodes
	nodes map[int64]graph.Node
	edges map[int64]map[int64]graph.Line
}

func NewLazyOrderedNodesByLines(nodes map[int64]graph.Node, edges map[int64]map[int64]graph.Line) *LazyOrderedNodesByLines {
	_ = "STUB: not implemented"
	return nil
}

func (n *LazyOrderedNodesByLines) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *LazyOrderedNodesByLines) Next() bool { _ = "STUB: not implemented"; return false }

func (n *LazyOrderedNodesByLines) Node() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (n *LazyOrderedNodesByLines) NodeSlice() []graph.Node { _ = "STUB: not implemented"; return nil }

func (n *LazyOrderedNodesByLines) Reset() { _ = "STUB: not implemented"; return }

func (n *LazyOrderedNodesByLines) fillSlice() { _ = "STUB: not implemented"; return }

type LazyOrderedNodesByWeightedLines struct {
	iter  OrderedNodes
	nodes map[int64]graph.Node
	edges map[int64]map[int64]graph.WeightedLine
}

func NewLazyOrderedNodesByWeightedLines(nodes map[int64]graph.Node, edges map[int64]map[int64]graph.WeightedLine) *LazyOrderedNodesByWeightedLines {
	_ = "STUB: not implemented"
	return nil
}

func (n *LazyOrderedNodesByWeightedLines) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *LazyOrderedNodesByWeightedLines) Next() bool { _ = "STUB: not implemented"; return false }

func (n *LazyOrderedNodesByWeightedLines) Node() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (n *LazyOrderedNodesByWeightedLines) NodeSlice() []graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (n *LazyOrderedNodesByWeightedLines) Reset() { _ = "STUB: not implemented"; return }

func (n *LazyOrderedNodesByWeightedLines) fillSlice() { _ = "STUB: not implemented"; return }

type ImplicitNodes struct {
	beg, end int
	curr     int
	newNode  func(id int) graph.Node
}

func NewImplicitNodes(beg, end int, new func(id int) graph.Node) *ImplicitNodes {
	_ = "STUB: not implemented"
	return nil
}

func (n *ImplicitNodes) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *ImplicitNodes) Next() bool { _ = "STUB: not implemented"; return false }

func (n *ImplicitNodes) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (n *ImplicitNodes) Reset() { _ = "STUB: not implemented"; return }

func (n *ImplicitNodes) NodeSlice() []graph.Node { _ = "STUB: not implemented"; return nil }
