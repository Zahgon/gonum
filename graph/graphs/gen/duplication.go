package gen

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
)

type UndirectedMutator interface {
	graph.UndirectedBuilder
	graph.EdgeRemover
}

func Duplication(dst UndirectedMutator, n int, delta, alpha, sigma float64, src rand.Source) error {
	_ = "STUB: not implemented"
	return nil
}
