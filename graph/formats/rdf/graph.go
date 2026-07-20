package rdf

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/set/uid"
)

type Graph struct {
	nodes map[int64]graph.Node
	from  map[int64]map[int64]map[int64]graph.Line
	to    map[int64]map[int64]map[int64]graph.Line
	pred  map[int64]map[*Statement]bool

	termIDs map[string]int64
	ids     *uid.Set
}

func NewGraph() *Graph { _ = "STUB: not implemented"; return nil }

func (g *Graph) addNode(n graph.Node) { _ = "STUB: not implemented"; return }

func (g *Graph) AddStatement(s *Statement) { _ = "STUB: not implemented"; return }

func (g *Graph) addTerm(t *Term) { _ = "STUB: not implemented"; return }

func (g *Graph) AllStatements() *Statements { _ = "STUB: not implemented"; return nil }

func (g *Graph) Edge(uid, vid int64) graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (g *Graph) Edges() graph.Edges { _ = "STUB: not implemented"; return *new(graph.Edges) }

func (g *Graph) From(id int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *Graph) FromSubject(t Term) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *Graph) HasEdgeBetween(xid, yid int64) bool { _ = "STUB: not implemented"; return false }

func (g *Graph) HasEdgeFromTo(uid, vid int64) bool { _ = "STUB: not implemented"; return false }

func (g *Graph) Lines(uid, vid int64) graph.Lines {
	_ = "STUB: not implemented"
	return *new(graph.Lines)
}

func (g *Graph) newLine(from, to graph.Node) graph.Line {
	_ = "STUB: not implemented"
	return *new(graph.Line)
}

func (g *Graph) newNode() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (g *Graph) Node(id int64) graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (g *Graph) TermFor(text string) (term Term, ok bool) {
	_ = "STUB: not implemented"
	return *new(Term), false
}

func (g *Graph) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *Graph) Predicates() []Term { _ = "STUB: not implemented"; return nil }

func (g *Graph) removeLine(fid, tid, id int64) { _ = "STUB: not implemented"; return }

func (g *Graph) removeNode(id int64) { _ = "STUB: not implemented"; return }

func (g *Graph) RemoveStatement(s *Statement) { _ = "STUB: not implemented"; return }

func (g *Graph) RemoveTerm(t Term) { _ = "STUB: not implemented"; return }

func (g *Graph) setLine(l graph.Line) { _ = "STUB: not implemented"; return }

func (g *Graph) Statements(uid, vid int64) *Statements { _ = "STUB: not implemented"; return nil }

func (g *Graph) To(id int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *Graph) ToObject(t Term) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

type Statements struct {
	eit graph.Edges
	lit graph.Lines
}

func (s *Statements) Next() bool { _ = "STUB: not implemented"; return false }

func (s *Statements) Statement() *Statement { _ = "STUB: not implemented"; return nil }

func ConnectedByAny(e graph.Edge, with func(*Statement) bool) bool {
	_ = "STUB: not implemented"
	return false
}
