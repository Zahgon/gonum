package graph

type Line interface {
	From() Node

	To() Node

	ReversedLine() Line

	ID() int64
}

type WeightedLine interface {
	Line
	Weight() float64
}

type Multigraph interface {
	Node(id int64) Node

	Nodes() Nodes

	From(id int64) Nodes

	HasEdgeBetween(xid, yid int64) bool

	Lines(uid, vid int64) Lines
}

type WeightedMultigraph interface {
	Multigraph

	WeightedLines(uid, vid int64) WeightedLines
}

type UndirectedMultigraph interface {
	Multigraph

	LinesBetween(xid, yid int64) Lines
}

type WeightedUndirectedMultigraph interface {
	WeightedMultigraph

	WeightedLinesBetween(xid, yid int64) WeightedLines
}

type DirectedMultigraph interface {
	Multigraph

	HasEdgeFromTo(uid, vid int64) bool

	To(id int64) Nodes
}

type WeightedDirectedMultigraph interface {
	WeightedMultigraph

	HasEdgeFromTo(uid, vid int64) bool

	To(id int64) Nodes
}

type LineAdder interface {
	NewLine(from, to Node) Line

	SetLine(l Line)
}

type WeightedLineAdder interface {
	NewWeightedLine(from, to Node, weight float64) WeightedLine

	SetWeightedLine(l WeightedLine)
}

type LineRemover interface {
	RemoveLine(fid, tid, id int64)
}

type MultigraphBuilder interface {
	NodeAdder
	LineAdder
}

type WeightedMultigraphBuilder interface {
	NodeAdder
	WeightedLineAdder
}

type UndirectedMultigraphBuilder interface {
	UndirectedMultigraph
	MultigraphBuilder
}

type UndirectedWeightedMultigraphBuilder interface {
	UndirectedMultigraph
	WeightedMultigraphBuilder
}

type DirectedMultigraphBuilder interface {
	DirectedMultigraph
	MultigraphBuilder
}

type DirectedWeightedMultigraphBuilder interface {
	DirectedMultigraph
	WeightedMultigraphBuilder
}
