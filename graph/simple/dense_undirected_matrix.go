package simple

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/mat"
)

var (
	um *UndirectedMatrix

	_ graph.Graph        = um
	_ graph.Undirected   = um
	_ edgeSetter         = um
	_ weightedEdgeSetter = um
)

type UndirectedMatrix struct {
	mat   *mat.SymDense
	nodes []graph.Node

	self   float64
	absent float64
}

func NewUndirectedMatrix(n int, init, self, absent float64) *UndirectedMatrix {
	_ = "STUB: not implemented"
	return nil
}

func NewUndirectedMatrixFrom(nodes []graph.Node, init, self, absent float64) *UndirectedMatrix {
	_ = "STUB: not implemented"
	return nil
}

func (g *UndirectedMatrix) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *UndirectedMatrix) EdgeBetween(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (g *UndirectedMatrix) Edges() graph.Edges { _ = "STUB: not implemented"; return *new(graph.Edges) }

func (g *UndirectedMatrix) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *UndirectedMatrix) HasEdgeBetween(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *UndirectedMatrix) Matrix() mat.Matrix { _ = "STUB: not implemented"; return *new(mat.Matrix) }

func (g *UndirectedMatrix) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (g *UndirectedMatrix) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *UndirectedMatrix) RemoveEdge(fid, tid int64) { _ = "STUB: not implemented"; return }

func (g *UndirectedMatrix) SetEdge(e graph.Edge) { _ = "STUB: not implemented"; return }

func (g *UndirectedMatrix) SetWeightedEdge(e graph.WeightedEdge) { _ = "STUB: not implemented"; return }

func (g *UndirectedMatrix) setWeightedEdge(e graph.Edge, weight float64) {
	_ = "STUB: not implemented"
	return
}

func (g *UndirectedMatrix) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (g *UndirectedMatrix) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *UndirectedMatrix) WeightedEdgeBetween(uid, vid int64) graph.WeightedEdge {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdge)
}

func (g *UndirectedMatrix) WeightedEdges() graph.WeightedEdges {
	_ = "STUB: not implemented"
	return *new(graph.WeightedEdges)
}

func (g *UndirectedMatrix) has(id int64) bool { _ = "STUB: not implemented"; return false }
