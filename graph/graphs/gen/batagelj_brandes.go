package gen

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
)

func Gnp(dst graph.Builder, n int, p float64, src rand.Source) error {
	_ = "STUB: not implemented"
	return nil
}

func Gnm(dst GraphBuilder, n, m int, src rand.Source) error { _ = "STUB: not implemented"; return nil }

func SmallWorldsBB(dst GraphBuilder, n, d int, p float64, src rand.Source) error {
	_ = "STUB: not implemented"
	return nil
}

func edgeNodesFor(i int, nodes []graph.Node) (v, w graph.Node) {
	_ = "STUB: not implemented"
	return *new(graph.Node), *new(graph.Node)
}

func PowerLaw(dst graph.MultigraphBuilder, n, d int, src rand.Source) error {
	_ = "STUB: not implemented"
	return nil
}

func BipartitePowerLaw(dst graph.MultigraphBuilder, n, d int, src rand.Source) (p1, p2 []graph.Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
