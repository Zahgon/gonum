package layout

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/spatial/r2"
)

type LayoutR2 interface {
	IsInitialized() bool

	SetCoord2(id int64, coords r2.Vec)

	Coord2(id int64) r2.Vec
}

func NewOptimizerR2(g graph.Graph, update func(graph.Graph, LayoutR2) bool) OptimizerR2 {
	_ = "STUB: not implemented"
	return *new(OptimizerR2)
}

type coordinatesR2 map[int64]r2.Vec

func (c coordinatesR2) IsInitialized() bool            { _ = "STUB: not implemented"; return false }
func (c coordinatesR2) SetCoord2(id int64, pos r2.Vec) { _ = "STUB: not implemented"; return }
func (c coordinatesR2) Coord2(id int64) r2.Vec         { _ = "STUB: not implemented"; return *new(r2.Vec) }

type OptimizerR2 struct {
	g      graph.Graph
	layout LayoutR2

	Updater func(graph.Graph, LayoutR2) bool
}

func (g OptimizerR2) Coord2(id int64) r2.Vec { _ = "STUB: not implemented"; return *new(r2.Vec) }

func (g OptimizerR2) Update() bool { _ = "STUB: not implemented"; return false }

func (g OptimizerR2) LayoutNodeR2(id int64) NodeR2 { _ = "STUB: not implemented"; return *new(NodeR2) }

func (g OptimizerR2) Node(id int64) graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (g OptimizerR2) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g OptimizerR2) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g OptimizerR2) HasEdgeBetween(xid, yid int64) bool { _ = "STUB: not implemented"; return false }

func (g OptimizerR2) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}
