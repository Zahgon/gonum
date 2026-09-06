package simple

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/set/uid"
)

var (
	ug *UndirectedGraph

	_ graph.Graph       = ug
	_ graph.Undirected  = ug
	_ graph.NodeAdder   = ug
	_ graph.NodeRemover = ug
	_ graph.EdgeAdder   = ug
	_ graph.EdgeRemover = ug
)

type UndirectedGraph struct {
	nodes map[int64]graph.Node
	edges map[int64]map[int64]graph.Edge

	nodeIDs *uid.Set
}

func NewUndirectedGraph() *UndirectedGraph { _ = "STUB: not implemented"; return nil }

func (g *UndirectedGraph) AddNode(n graph.Node) { _ = "STUB: not implemented"; return }

func (g *UndirectedGraph) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *UndirectedGraph) EdgeBetween(xid, yid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *UndirectedGraph) Edges() graph.Edges { _ = "STUB: not implemented"; return *new(graph.Edges) }

func (g *UndirectedGraph) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *UndirectedGraph) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *UndirectedGraph) NewEdge(from, to graph.Node) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *UndirectedGraph) NewNode() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (g *UndirectedGraph) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *UndirectedGraph) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *UndirectedGraph) NodeWithID(id int64) (n graph.Node, new bool) {
	_ = "STUB: not implemented"
	return *new(graph.Node), false
}

func (g *UndirectedGraph) RemoveEdge(fid, tid int64) { _ = "STUB: not implemented"; return }

func (g *UndirectedGraph) RemoveNode(id int64) { _ = "STUB: not implemented"; return }

func (g *UndirectedGraph) SetEdge(e graph.Edge) { _ = "STUB: not implemented"; return }
