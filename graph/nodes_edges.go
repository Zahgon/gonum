package graph

type Iterator interface {
	Next() bool

	Len() int

	Reset()
}

type Nodes interface {
	Iterator

	Node() Node
}

type NodeSlicer interface {
	NodeSlice() []Node
}

func NodesOf(it Nodes) []Node { _ = "STUB: not implemented"; return nil }

type Edges interface {
	Iterator

	Edge() Edge
}

type EdgeSlicer interface {
	EdgeSlice() []Edge
}

func EdgesOf(it Edges) []Edge { _ = "STUB: not implemented"; return nil }

type WeightedEdges interface {
	Iterator

	WeightedEdge() WeightedEdge
}

type WeightedEdgeSlicer interface {
	WeightedEdgeSlice() []WeightedEdge
}

func WeightedEdgesOf(it WeightedEdges) []WeightedEdge { _ = "STUB: not implemented"; return nil }

type Lines interface {
	Iterator

	Line() Line
}

type LineSlicer interface {
	LineSlice() []Line
}

func LinesOf(it Lines) []Line { _ = "STUB: not implemented"; return nil }

type WeightedLines interface {
	Iterator

	WeightedLine() WeightedLine
}

type WeightedLineSlicer interface {
	WeightedLineSlice() []WeightedLine
}

func WeightedLinesOf(it WeightedLines) []WeightedLine { _ = "STUB: not implemented"; return nil }

const Empty = nothing

var (
	_ Iterator           = Empty
	_ Nodes              = Empty
	_ NodeSlicer         = Empty
	_ Edges              = Empty
	_ EdgeSlicer         = Empty
	_ WeightedEdges      = Empty
	_ WeightedEdgeSlicer = Empty
	_ Lines              = Empty
	_ LineSlicer         = Empty
	_ WeightedLines      = Empty
	_ WeightedLineSlicer = Empty
)

const nothing = empty(0)

type empty int

func (empty) Next() bool                        { _ = "STUB: not implemented"; return false }
func (empty) Len() int                          { _ = "STUB: not implemented"; return 0 }
func (empty) Reset()                            { _ = "STUB: not implemented"; return }
func (empty) Node() Node                        { _ = "STUB: not implemented"; return *new(Node) }
func (empty) NodeSlice() []Node                 { _ = "STUB: not implemented"; return nil }
func (empty) Edge() Edge                        { _ = "STUB: not implemented"; return *new(Edge) }
func (empty) EdgeSlice() []Edge                 { _ = "STUB: not implemented"; return nil }
func (empty) WeightedEdge() WeightedEdge        { _ = "STUB: not implemented"; return *new(WeightedEdge) }
func (empty) WeightedEdgeSlice() []WeightedEdge { _ = "STUB: not implemented"; return nil }
func (empty) Line() Line                        { _ = "STUB: not implemented"; return *new(Line) }
func (empty) LineSlice() []Line                 { _ = "STUB: not implemented"; return nil }
func (empty) WeightedLine() WeightedLine        { _ = "STUB: not implemented"; return *new(WeightedLine) }
func (empty) WeightedLineSlice() []WeightedLine { _ = "STUB: not implemented"; return nil }

func (empty) String() string   { _ = "STUB: not implemented"; return "" }
func (empty) GoString() string { _ = "STUB: not implemented"; return "" }
