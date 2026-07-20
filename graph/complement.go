package graph

type Complement struct {
	Graph
}

func (g Complement) Edge(uid, vid int64) Edge { _ = "STUB: not implemented"; return *new(Edge) }

func (g Complement) From(uid int64) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (g Complement) HasEdgeBetween(xid, yid int64) bool { _ = "STUB: not implemented"; return false }

type shadow struct{ F, T Node }

func (e shadow) From() Node         { _ = "STUB: not implemented"; return *new(Node) }
func (e shadow) To() Node           { _ = "STUB: not implemented"; return *new(Node) }
func (e shadow) ReversedEdge() Edge { _ = "STUB: not implemented"; return *new(Edge) }

type nodeFilterIterator struct {
	src Nodes

	filter map[int64]bool
}

func newNodeFilterIterator(src, filter Nodes, root int64) *nodeFilterIterator {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeFilterIterator) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *nodeFilterIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (n *nodeFilterIterator) Node() Node { _ = "STUB: not implemented"; return *new(Node) }

func (n *nodeFilterIterator) Reset() { _ = "STUB: not implemented"; return }
