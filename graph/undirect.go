package graph

type Undirect struct {
	G Directed
}

var _ Undirected = Undirect{}

func (g Undirect) Node(id int64) Node { _ = "STUB: not implemented"; return *new(Node) }

func (g Undirect) Nodes() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (g Undirect) From(uid int64) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (g Undirect) HasEdgeBetween(xid, yid int64) bool { _ = "STUB: not implemented"; return false }

func (g Undirect) Edge(uid, vid int64) Edge { _ = "STUB: not implemented"; return *new(Edge) }

func (g Undirect) EdgeBetween(xid, yid int64) Edge { _ = "STUB: not implemented"; return *new(Edge) }

type UndirectWeighted struct {
	G WeightedDirected

	Absent float64

	Merge func(x, y float64, xe, ye Edge) float64
}

var (
	_ Undirected         = UndirectWeighted{}
	_ WeightedUndirected = UndirectWeighted{}
)

func (g UndirectWeighted) Node(id int64) Node { _ = "STUB: not implemented"; return *new(Node) }

func (g UndirectWeighted) Nodes() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (g UndirectWeighted) From(uid int64) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (g UndirectWeighted) HasEdgeBetween(xid, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (g UndirectWeighted) Edge(uid, vid int64) Edge { _ = "STUB: not implemented"; return *new(Edge) }

func (g UndirectWeighted) WeightedEdge(uid, vid int64) WeightedEdge {
	_ = "STUB: not implemented"
	return *new(WeightedEdge)
}

func (g UndirectWeighted) EdgeBetween(xid, yid int64) Edge {
	_ = "STUB: not implemented"
	return *new(Edge)
}

func (g UndirectWeighted) WeightedEdgeBetween(xid, yid int64) WeightedEdge {
	_ = "STUB: not implemented"
	return *new(WeightedEdge)
}

func (g UndirectWeighted) Weight(xid, yid int64) (w float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type EdgePair [2]Edge

func (e EdgePair) From() Node { _ = "STUB: not implemented"; return *new(Node) }

func (e EdgePair) To() Node { _ = "STUB: not implemented"; return *new(Node) }

func (e EdgePair) ReversedEdge() Edge { _ = "STUB: not implemented"; return *new(Edge) }

type WeightedEdgePair struct {
	EdgePair
	W float64
}

func (e WeightedEdgePair) ReversedEdge() Edge { _ = "STUB: not implemented"; return *new(Edge) }

func (e WeightedEdgePair) Weight() float64 { _ = "STUB: not implemented"; return 0 }

type nodeIteratorPair struct {
	a, b Nodes

	curr Node

	idx, cnt int

	unique map[int64]bool
}

func newNodeIteratorPair(a, b Nodes) *nodeIteratorPair { _ = "STUB: not implemented"; return nil }

func (n *nodeIteratorPair) Len() int { _ = "STUB: not implemented"; return 0 }

func (n *nodeIteratorPair) Next() bool { _ = "STUB: not implemented"; return false }

func (n *nodeIteratorPair) Node() Node { _ = "STUB: not implemented"; return *new(Node) }

func (n *nodeIteratorPair) Reset() { _ = "STUB: not implemented"; return }
