package path

type djSet map[int64]*dsNode

func (s djSet) add(e int64) { _ = "STUB: not implemented"; return }

func (djSet) union(a, b *dsNode) { _ = "STUB: not implemented"; return }

func (s djSet) find(e int64) *dsNode { _ = "STUB: not implemented"; return nil }

func find(n *dsNode) *dsNode { _ = "STUB: not implemented"; return nil }

type dsNode struct {
	parent *dsNode
	rank   int
}
