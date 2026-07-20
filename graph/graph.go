package graph

type Node interface {
	ID() int64
}

type Edge interface {
	From() Node

	To() Node

	ReversedEdge() Edge
}

type WeightedEdge interface {
	Edge
	Weight() float64
}

type Graph interface {
	Node(id int64) Node

	Nodes() Nodes

	From(id int64) Nodes

	HasEdgeBetween(xid, yid int64) bool

	Edge(uid, vid int64) Edge
}

type Weighted interface {
	Graph

	WeightedEdge(uid, vid int64) WeightedEdge

	Weight(xid, yid int64) (w float64, ok bool)
}

type Undirected interface {
	Graph

	EdgeBetween(xid, yid int64) Edge
}

type WeightedUndirected interface {
	Weighted

	WeightedEdgeBetween(xid, yid int64) WeightedEdge
}

type Directed interface {
	Graph

	HasEdgeFromTo(uid, vid int64) bool

	To(id int64) Nodes
}

type WeightedDirected interface {
	Weighted

	HasEdgeFromTo(uid, vid int64) bool

	To(id int64) Nodes
}

type NodeAdder interface {
	NewNode() Node

	AddNode(Node)
}

type NodeWithIDer interface {
	NodeWithID(id int64) (n Node, new bool)
}

type NodeRemover interface {
	RemoveNode(id int64)
}

type EdgeAdder interface {
	NewEdge(from, to Node) Edge

	SetEdge(e Edge)
}

type WeightedEdgeAdder interface {
	NewWeightedEdge(from, to Node, weight float64) WeightedEdge

	SetWeightedEdge(e WeightedEdge)
}

type EdgeRemover interface {
	RemoveEdge(fid, tid int64)
}

type Builder interface {
	NodeAdder
	EdgeAdder
}

type WeightedBuilder interface {
	NodeAdder
	WeightedEdgeAdder
}

type UndirectedBuilder interface {
	Undirected
	Builder
}

type UndirectedWeightedBuilder interface {
	Undirected
	WeightedBuilder
}

type DirectedBuilder interface {
	Directed
	Builder
}

type DirectedWeightedBuilder interface {
	Directed
	WeightedBuilder
}

func Copy(dst Builder, src Graph) { _ = "STUB: not implemented"; return }

func CopyWeighted(dst WeightedBuilder, src Weighted) { _ = "STUB: not implemented"; return }
