package rdf

import (
	"gonum.org/v1/gonum/graph"
)

type Query struct {
	g graph.Directed

	terms []Term
}

func NewQuery(g graph.Directed, from ...Term) Query { _ = "STUB: not implemented"; return *new(Query) }

func (g *Graph) Query(from ...Term) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q Query) Out(fn func(s *Statement) bool) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q Query) In(fn func(s *Statement) bool) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q Query) HasAllOut(fn func(s *Statement) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

func (q Query) HasAllIn(fn func(s *Statement) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

func (q Query) HasAnyOut(fn func(s *Statement) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

func (q Query) HasAnyIn(fn func(s *Statement) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

func not(fn func(s *Statement) bool) func(s *Statement) bool { _ = "STUB: not implemented"; return nil }

func (q Query) And(p Query) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q Query) Or(p Query) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q Query) Not(p Query) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q Query) Repeat(fn func(Query) (q Query, ok bool)) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

func (q Query) Unique() Query { _ = "STUB: not implemented"; return *new(Query) }

func (q Query) Len() int { _ = "STUB: not implemented"; return 0 }

func (q Query) Result() []Term { _ = "STUB: not implemented"; return nil }
