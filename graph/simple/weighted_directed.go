package simple

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/set/uid"
)

var (
	wdg *WeightedDirectedGraph

	_ graph.Graph             = wdg
	_ graph.Weighted          = wdg
	_ graph.Directed          = wdg
	_ graph.WeightedDirected  = wdg
	_ graph.NodeAdder         = wdg
	_ graph.NodeRemover       = wdg
	_ graph.WeightedEdgeAdder = wdg
	_ graph.EdgeRemover       = wdg
)

type WeightedDirectedGraph struct {
	nodes map[int64]graph.Node
	from  map[int64]map[int64]graph.WeightedEdge
	to    map[int64]map[int64]graph.WeightedEdge

	self, absent float64

	nodeIDs *uid.Set
}

func NewWeightedDirectedGraph(self, absent float64) *WeightedDirectedGraph {
	_ = "STUB: not implemented"
	return nil
}

func (g *WeightedDirectedGraph) AddNode(n graph.Node) { _ = "STUB: not implemented"; return }

func (g *WeightedDirectedGraph) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *WeightedDirectedGraph) Edges() graph.Edges {
	_ = "STUB: not implemented"
	return *new(graph.Edges)
}

func (g *WeightedDirectedGraph) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *WeightedDirectedGraph) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *WeightedDirectedGraph) HasEdgeFromTo(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *WeightedDirectedGraph) NewNode() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *WeightedDirectedGraph) NewWeightedEdge(from, to graph.Node, weight float64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *WeightedDirectedGraph) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *WeightedDirectedGraph) Nodes() graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *WeightedDirectedGraph) NodeWithID(id int64) (n graph.Node, new bool) {
	_ = "STUB: not implemented"
	return *new(graph.Node), false
}

func (g *WeightedDirectedGraph) RemoveEdge(fid, tid int64) { _ = "STUB: not implemented"; return }

func (g *WeightedDirectedGraph) RemoveNode(id int64) { _ = "STUB: not implemented"; return }

func (g *WeightedDirectedGraph) SetWeightedEdge(e graph.WeightedEdge) {
	_ = "STUB: not implemented"
	return
}

func (g *WeightedDirectedGraph) To(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *WeightedDirectedGraph) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (g *WeightedDirectedGraph) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *WeightedDirectedGraph) WeightedEdges() graph.WeightedEdges {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdges)
}
