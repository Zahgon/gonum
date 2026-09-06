package simple

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/set/uid"
)

var (
	dg *DirectedGraph

	_ graph.Graph       = dg
	_ graph.Directed    = dg
	_ graph.NodeAdder   = dg
	_ graph.NodeRemover = dg
	_ graph.EdgeAdder   = dg
	_ graph.EdgeRemover = dg
)

type DirectedGraph struct {
	nodes map[int64]graph.Node
	from  map[int64]map[int64]graph.Edge
	to    map[int64]map[int64]graph.Edge

	nodeIDs *uid.Set
}

func NewDirectedGraph() *DirectedGraph { _ = "STUB: not implemented"; return nil }

func (g *DirectedGraph) AddNode(n graph.Node) { _ = "STUB: not implemented"; return }

func (g *DirectedGraph) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *DirectedGraph) Edges() graph.Edges { _ = "STUB: not implemented"; return *new(graph.Edges) }

func (g *DirectedGraph) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *DirectedGraph) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *DirectedGraph) HasEdgeFromTo(uid, vid int64) bool { _ = "STUB: not implemented"; return false }

func (g *DirectedGraph) NewEdge(from, to graph.Node) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *DirectedGraph) NewNode() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (g *DirectedGraph) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *DirectedGraph) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *DirectedGraph) NodeWithID(id int64) (n graph.Node, new bool) {
	_ = "STUB: not implemented"
	return *new(graph.Node), false
}

func (g *DirectedGraph) RemoveEdge(fid, tid int64) { _ = "STUB: not implemented"; return }

func (g *DirectedGraph) RemoveNode(id int64) { _ = "STUB: not implemented"; return }

func (g *DirectedGraph) SetEdge(e graph.Edge) { _ = "STUB: not implemented"; return }

func (g *DirectedGraph) To(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}
