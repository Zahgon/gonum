package layout

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/spatial/barneshut"
	"gonum.org/v1/gonum/spatial/r2"
)

type EadesR2 struct {
	Updates int

	Repulsion float64

	Rate float64

	Theta float64

	Src rand.Source

	nodes   graph.Nodes
	indexOf map[int64]int

	particles []barneshut.Particle2
	forces    []r2.Vec
}

func (u *EadesR2) Update(g graph.Graph, layout LayoutR2) bool {
	_ = "STUB: not implemented"
	return false
}

type eadesR2Node struct {
	id  int64
	pos r2.Vec
}

func (p eadesR2Node) Coord2() r2.Vec { _ = "STUB: not implemented"; return *new(r2.Vec) }
func (p eadesR2Node) Mass() float64  { _ = "STUB: not implemented"; return 0 }
