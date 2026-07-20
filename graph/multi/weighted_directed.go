package multi

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/set/uid"
)

var (
	wdg *WeightedDirectedGraph

	_ graph.Graph                      = wdg
	_ graph.Weighted                   = wdg
	_ graph.Directed                   = wdg
	_ graph.WeightedDirected           = wdg
	_ graph.Multigraph                 = wdg
	_ graph.DirectedMultigraph         = wdg
	_ graph.WeightedDirectedMultigraph = wdg
	_ graph.NodeAdder                  = wdg
	_ graph.NodeRemover                = wdg
	_ graph.WeightedLineAdder          = wdg
	_ graph.LineRemover                = wdg
)

type WeightedDirectedGraph struct {
	EdgeWeightFunc func(graph.WeightedLines) float64

	nodes map[int64]graph.Node
	from  map[int64]map[int64]map[int64]graph.WeightedLine
	to    map[int64]map[int64]map[int64]graph.WeightedLine

	nodeIDs *uid.Set
	lineIDs map[int64]map[int64]*uid.Set
}

func NewWeightedDirectedGraph() *WeightedDirectedGraph { _ = "STUB: not implemented"; return nil }

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

func (g *WeightedDirectedGraph) Lines(uid, vid int64) graph.Lines {
	_ = "STUB: not implemented"
	return *new(graph.Lines)
}

func (g *WeightedDirectedGraph) NewNode() graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *WeightedDirectedGraph) NewWeightedLine(from, to graph.Node, weight float64) graph.WeightedLine {
	_ = "STUB: not implemented"
	return *new(graph.WeightedLine)
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

func (g *WeightedDirectedGraph) RemoveLine(fid, tid, id int64) { _ = "STUB: not implemented"; return }

func (g *WeightedDirectedGraph) RemoveNode(id int64) { _ = "STUB: not implemented"; return }

func (g *WeightedDirectedGraph) SetWeightedLine(l graph.WeightedLine) {
	_ = "STUB: not implemented"
	return
}

func (g *WeightedDirectedGraph) To(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *WeightedDirectedGraph) Weight(uid, vid int64) (w float64, ok bool) {
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

func (g *WeightedDirectedGraph) WeightedLines(uid, vid int64) graph.WeightedLines {
	_ = "STUB: not implemented"
	return *new(graph.WeightedLines)
}
