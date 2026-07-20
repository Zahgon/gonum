package encoding

import "gonum.org/v1/gonum/graph"

type Builder interface {
	graph.Graph
	graph.Builder
}

type MultiBuilder interface {
	graph.Multigraph
	graph.MultigraphBuilder
}

type AttributeSetter interface {
	SetAttribute(Attribute) error
}

type Attributer interface {
	Attributes() []Attribute
}

type Attribute struct {
	Key, Value string
}

type Attributes []Attribute

func (a *Attributes) Attributes() []Attribute { _ = "STUB: not implemented"; return nil }

func (a *Attributes) SetAttribute(attr Attribute) error { _ = "STUB: not implemented"; return nil }
