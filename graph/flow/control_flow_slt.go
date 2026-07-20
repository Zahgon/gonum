package flow

import "gonum.org/v1/gonum/graph"

func DominatorsSLT(root graph.Node, g graph.Directed) DominatorTree {
	_ = "STUB: not implemented"
	return *new(DominatorTree)
}

type sLengauerTarjan struct {
	nodes []*sltNode

	indexOf map[int64]int

	base sltNode
}

type sltNode struct {
	node graph.Node

	parent *sltNode

	pred []*sltNode

	semi int

	size int

	child *sltNode

	bucket map[*sltNode]struct{}

	dom *sltNode

	ancestor *sltNode

	label *sltNode
}

func (lt *sLengauerTarjan) dfs(g graph.Directed, v graph.Node) { _ = "STUB: not implemented"; return }

func (lt *sLengauerTarjan) compress(v *sltNode) { _ = "STUB: not implemented"; return }

func (lt *sLengauerTarjan) eval(v *sltNode) *sltNode { _ = "STUB: not implemented"; return nil }

func (*sLengauerTarjan) link(v, w *sltNode) { _ = "STUB: not implemented"; return }
