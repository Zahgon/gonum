package digraph6

import (
	"math/big"

	"gonum.org/v1/gonum/graph"
)

type Graph string

var (
	d6 Graph

	_ graph.Graph    = d6
	_ graph.Directed = d6
)

func Encode(g graph.Graph) Graph { _ = "STUB: not implemented"; return *new(Graph) }

func bit6(b int64) byte { _ = "STUB: not implemented"; return 0 }

func IsValid(g Graph) bool { _ = "STUB: not implemented"; return false }

func (g Graph) Edge(uid, vid int64) graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (g Graph) From(id int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g Graph) HasEdgeBetween(xid, yid int64) bool { _ = "STUB: not implemented"; return false }

func (g Graph) HasEdgeFromTo(uid, vid int64) bool { _ = "STUB: not implemented"; return false }

func (g Graph) Node(id int64) graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (g Graph) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g Graph) To(id int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

type d6ForwardIterator struct {
	g    Graph
	n    int64
	from int64
	to   int64
}

var _ graph.Nodes = (*d6ForwardIterator)(nil)

func (i *d6ForwardIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (i *d6ForwardIterator) Len() int { _ = "STUB: not implemented"; return 0 }

func (i *d6ForwardIterator) Reset() { _ = "STUB: not implemented"; return }

func (i *d6ForwardIterator) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

type d6ReverseIterator struct {
	g    Graph
	n    int64
	from int64
	to   int64
}

var _ graph.Nodes = (*d6ReverseIterator)(nil)

func (i *d6ReverseIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (i *d6ReverseIterator) Len() int { _ = "STUB: not implemented"; return 0 }

func (i *d6ReverseIterator) Reset() { _ = "STUB: not implemented"; return }

func (i *d6ReverseIterator) Node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func numberOf(g Graph) int64 { _ = "STUB: not implemented"; return 0 }

func bitFor(uid, vid, n int64) int { _ = "STUB: not implemented"; return 0 }

func isSet(bit int, g Graph) bool { _ = "STUB: not implemented"; return false }

func (g Graph) GoString() string { _ = "STUB: not implemented"; return "" }

func binary(g Graph) (b *big.Int, l int) { _ = "STUB: not implemented"; return nil, 0 }
