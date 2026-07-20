package flow

import "gonum.org/v1/gonum/graph"

func Dominators(root graph.Node, g graph.Directed) DominatorTree {
	_ = "STUB: not implemented"
	return *new(DominatorTree)
}

type lengauerTarjan struct {
	nodes []*ltNode

	indexOf map[int64]int
}

type ltNode struct {
	node graph.Node

	parent *ltNode

	pred []*ltNode

	semi int

	bucket map[*ltNode]struct{}

	dom *ltNode

	ancestor *ltNode

	label *ltNode
}

func (lt *lengauerTarjan) dfs(g graph.Directed, v graph.Node) { _ = "STUB: not implemented"; return }

func (lt *lengauerTarjan) compress(v *ltNode) { _ = "STUB: not implemented"; return }

func (lt *lengauerTarjan) eval(v *ltNode) *ltNode { _ = "STUB: not implemented"; return nil }

func (*lengauerTarjan) link(v, w *ltNode) { _ = "STUB: not implemented"; return }

type DominatorTree struct {
	root        graph.Node
	dominatorOf map[int64]graph.Node
	dominatedBy map[int64][]graph.Node
}

func (d DominatorTree) Root() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (d DominatorTree) DominatorOf(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (d DominatorTree) DominatedBy(id int64) []graph.Node { _ = "STUB: not implemented"; return nil }
