package ast

import (
	"fmt"
)

type File struct {
	Graphs []*Graph
}

func (f *File) String() string { _ = "STUB: not implemented"; return "" }

type Graph struct {
	Strict bool

	Directed bool

	ID string

	Stmts []Stmt
}

func (g *Graph) String() string { _ = "STUB: not implemented"; return "" }

type Stmt interface {
	fmt.Stringer

	isStmt()
}

type NodeStmt struct {
	Node *Node

	Attrs []*Attr
}

func (e *NodeStmt) String() string { _ = "STUB: not implemented"; return "" }

type EdgeStmt struct {
	From Vertex

	To *Edge

	Attrs []*Attr
}

func (e *EdgeStmt) String() string { _ = "STUB: not implemented"; return "" }

type Edge struct {
	Directed bool

	Vertex Vertex

	To *Edge
}

func (e *Edge) String() string { _ = "STUB: not implemented"; return "" }

type AttrStmt struct {
	Kind Kind

	Attrs []*Attr
}

func (a *AttrStmt) String() string { _ = "STUB: not implemented"; return "" }

type Kind uint

const (
	GraphKind Kind = iota
	NodeKind
	EdgeKind
)

func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

type Attr struct {
	Key string

	Val string
}

func (a *Attr) String() string { _ = "STUB: not implemented"; return "" }

type Subgraph struct {
	ID string

	Stmts []Stmt
}

func (s *Subgraph) String() string { _ = "STUB: not implemented"; return "" }

func (*NodeStmt) isStmt() { _ = "STUB: not implemented"; return }
func (*EdgeStmt) isStmt() { _ = "STUB: not implemented"; return }
func (*AttrStmt) isStmt() { _ = "STUB: not implemented"; return }
func (*Attr) isStmt()     { _ = "STUB: not implemented"; return }
func (*Subgraph) isStmt() { _ = "STUB: not implemented"; return }

type Vertex interface {
	fmt.Stringer

	isVertex()
}

type Node struct {
	ID string

	Port *Port
}

func (n *Node) String() string { _ = "STUB: not implemented"; return "" }

type Port struct {
	ID string

	CompassPoint CompassPoint
}

func (p *Port) String() string { _ = "STUB: not implemented"; return "" }

type CompassPoint uint

const (
	CompassPointNone CompassPoint = iota
	CompassPointNorth
	CompassPointNorthEast
	CompassPointEast
	CompassPointSouthEast
	CompassPointSouth
	CompassPointSouthWest
	CompassPointWest
	CompassPointNorthWest
	CompassPointCenter
	CompassPointDefault
)

func (c CompassPoint) String() string { _ = "STUB: not implemented"; return "" }

func (*Node) isVertex()     { _ = "STUB: not implemented"; return }
func (*Subgraph) isVertex() { _ = "STUB: not implemented"; return }
