package rdf

type djSet map[string]*dsNode

func (s djSet) add(e string) { _ = "STUB: not implemented"; return }

func (djSet) union(a, b *dsNode) { _ = "STUB: not implemented"; return }

func (s djSet) find(e string) *dsNode { _ = "STUB: not implemented"; return nil }

func find(n *dsNode) *dsNode { _ = "STUB: not implemented"; return nil }

type dsNode struct {
	parent *dsNode
	rank   int
}
