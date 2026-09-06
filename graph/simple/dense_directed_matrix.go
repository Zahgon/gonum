package simple

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/mat"
)

var (
	dm *DirectedMatrix

	_ graph.Graph        = dm
	_ graph.Directed     = dm
	_ edgeSetter         = dm
	_ weightedEdgeSetter = dm
)

type DirectedMatrix struct {
	mat   *mat.Dense
	nodes []graph.Node

	self   float64
	absent float64
}

func NewDirectedMatrix(n int, init, self, absent float64) *DirectedMatrix {
	_ = "STUB: not implemented"
	return nil
}

func NewDirectedMatrixFrom(nodes []graph.Node, init, self, absent float64) *DirectedMatrix {
	_ = "STUB: not implemented"
	return nil
}

func (g *DirectedMatrix) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *DirectedMatrix) Edges() graph.Edges { _ = "STUB: not implemented"; return *new(graph.Edges) }

func (g *DirectedMatrix) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *DirectedMatrix) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *DirectedMatrix) HasEdgeFromTo(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *DirectedMatrix) Matrix() mat.Matrix { _ = "STUB: not implemented"; return *new(mat.Matrix) }

func (g *DirectedMatrix) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *DirectedMatrix) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *DirectedMatrix) RemoveEdge(fid, tid int64) { _ = "STUB: not implemented"; return }

func (g *DirectedMatrix) SetEdge(e graph.Edge) { _ = "STUB: not implemented"; return }

func (g *DirectedMatrix) SetWeightedEdge(e graph.WeightedEdge) { _ = "STUB: not implemented"; return }

func (g *DirectedMatrix) setWeightedEdge(e graph.Edge, weight float64) {
	_ = "STUB: not implemented"
	return
}

func (g *DirectedMatrix) To(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *DirectedMatrix) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (g *DirectedMatrix) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *DirectedMatrix) WeightedEdges() graph.WeightedEdges {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdges)
}

func (g *DirectedMatrix) has(id int64) bool { _ = "STUB: not implemented"; return false }
