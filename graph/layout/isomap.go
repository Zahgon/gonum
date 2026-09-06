package layout

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/mat"
)

type IsomapR2 struct{}

func (IsomapR2) Update(g graph.Graph, layout LayoutR2) bool {
	_ = "STUB: not implemented"
	return false
}

func isomap(g graph.Graph, nodes []graph.Node, dims int) *mat.Dense {
	_ = "STUB: not implemented"
	return nil
}
