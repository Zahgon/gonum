package simple

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/set/uid"
)

var (
	wug *WeightedUndirectedGraph

	_ graph.Graph              = wug
	_ graph.Weighted           = wug
	_ graph.Undirected         = wug
	_ graph.WeightedUndirected = wug
	_ graph.NodeAdder          = wug
	_ graph.NodeRemover        = wug
	_ graph.WeightedEdgeAdder  = wug
	_ graph.EdgeRemover        = wug
)

type WeightedUndirectedGraph struct {
	nodes map[int64]graph.Node
	edges map[int64]map[int64]graph.WeightedEdge

	self, absent float64

	nodeIDs *uid.Set
}

func NewWeightedUndirectedGraph(self, absent float64) *WeightedUndirectedGraph {
	_ = "STUB: not implemented"
	return nil
}

func (g *WeightedUndirectedGraph) AddNode(n graph.Node) { _ = "STUB: not implemented"; return }

func (g *WeightedUndirectedGraph) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *WeightedUndirectedGraph) EdgeBetween(xid, yid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *WeightedUndirectedGraph) Edges() graph.Edges {
	_ = "STUB: not implemented"
	return *new(graph.Edges)
}

func (g *WeightedUndirectedGraph) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *WeightedUndirectedGraph) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *WeightedUndirectedGraph) NewNode() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *WeightedUndirectedGraph) NewWeightedEdge(from, to graph.Node, weight float64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *WeightedUndirectedGraph) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *WeightedUndirectedGraph) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *WeightedUndirectedGraph) NodeWithID(id int64) (n graph.Node, new bool) {
	_ = "STUB: not implemented"
	return *new(graph.Node), false
}

func (g *WeightedUndirectedGraph) RemoveEdge(fid, tid int64) { _ = "STUB: not implemented"; return }

func (g *WeightedUndirectedGraph) RemoveNode(id int64) { _ = "STUB: not implemented"; return }

func (g *WeightedUndirectedGraph) SetWeightedEdge(e graph.WeightedEdge) {
	_ = "STUB: not implemented"
	return
}

func (g *WeightedUndirectedGraph) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (g *WeightedUndirectedGraph) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *WeightedUndirectedGraph) WeightedEdgeBetween(xid, yid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *WeightedUndirectedGraph) WeightedEdges() graph.WeightedEdges {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdges)
}
