package layout

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/spatial/r2"
)

type GraphR2 interface {
	graph.Graph
	LayoutNodeR2(id int64) NodeR2
}

type NodeR2 struct {
	graph.Node
	Coord2 r2.Vec
}
