package spectral

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/mat"
)

type Laplacian struct {
	mat.Matrix

	Nodes []graph.Node

	Index map[int64]int
}

func NewLaplacian(g graph.Undirected) Laplacian { _ = "STUB: not implemented"; return *new(Laplacian) }

func NewSymNormLaplacian(g graph.Undirected) Laplacian {
	_ = "STUB: not implemented"
	return *new(Laplacian)
}

func NewRandomWalkLaplacian(g graph.Graph, damp float64) Laplacian {
	_ = "STUB: not implemented"
	return *new(Laplacian)
}
