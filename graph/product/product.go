package product

import (
	"gonum.org/v1/gonum/graph"
)

type Node struct {
	UID int64

	A, B graph.Node
}

func (n Node) ID() int64 { _ = "STUB: not implemented"; return 0 }

func Cartesian(dst graph.Builder, a, b graph.Graph) { _ = "STUB: not implemented"; return }

func Tensor(dst graph.Builder, a, b graph.Graph) { _ = "STUB: not implemented"; return }

func Lexicographical(dst graph.Builder, a, b graph.Graph) { _ = "STUB: not implemented"; return }

func Strong(dst graph.Builder, a, b graph.Graph) { _ = "STUB: not implemented"; return }

func CoNormal(dst graph.Builder, a, b graph.Graph) { _ = "STUB: not implemented"; return }

func Modular(dst graph.Builder, a, b graph.Graph) { _ = "STUB: not implemented"; return }

func ModularExt(dst graph.Builder, a, b graph.Graph, agree func(eA, eB graph.Edge) bool) {
	_ = "STUB: not implemented"
	return
}

func cartesianNodes(a, b graph.Graph) (aNodes, bNodes []graph.Node, product []Node) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func lexicalNodes(g graph.Graph) []graph.Node { _ = "STUB: not implemented"; return nil }

func indexOf(nodes []graph.Node) map[int64]int { _ = "STUB: not implemented"; return nil }
