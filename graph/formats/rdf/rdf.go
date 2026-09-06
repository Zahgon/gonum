//go:generate ragel -Z -G2 parse.rl
//go:generate ragel -Z -G2 extract.rl
//go:generate ragel -Z -G2 check.rl
//go:generate go tool golang.org/x/tools/cmd/stringer -type=Kind

package rdf

import (
	"bufio"
	"errors"
	"io"

	"gonum.org/v1/gonum/graph"
)

var (
	_ graph.Node = Term{}
	_ graph.Edge = (*Statement)(nil)
	_ graph.Line = (*Statement)(nil)
)

var (
	ErrInvalid        = errors.New("invalid N-Quad")
	ErrIncomplete     = errors.New("incomplete N-Quad")
	ErrInvalidTerm    = errors.New("invalid term")
	ErrIncompleteTerm = errors.New("incomplete term")
)

type Kind int

const (
	Invalid Kind = iota

	IRI

	Literal

	Blank
)

type Term struct {
	Value string

	UID int64
}

func NewBlankTerm(label string) (Term, error) { _ = "STUB: not implemented"; return *new(Term), nil }

const blankPrefix = "_:"

func isBlank(s string) bool { _ = "STUB: not implemented"; return false }

func NewIRITerm(iri string) (Term, error) { _ = "STUB: not implemented"; return *new(Term), nil }

func isIRI(s string) bool { _ = "STUB: not implemented"; return false }

func NewLiteralTerm(text, qual string) (Term, error) {
	_ = "STUB: not implemented"
	return *new(Term), nil
}

func checkIRIText(iri string) error { _ = "STUB: not implemented"; return nil }

func isLiteral(s string) bool { _ = "STUB: not implemented"; return false }

func (t Term) Parts() (text, qual string, kind Kind, err error) {
	_ = "STUB: not implemented"
	return "", "", *new(Kind), nil
}

func (t Term) ID() int64 { _ = "STUB: not implemented"; return 0 }

type Statement struct {
	Subject   Term
	Predicate Term
	Object    Term
	Label     Term
}

func (s *Statement) String() string { _ = "STUB: not implemented"; return "" }

func (s *Statement) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (s *Statement) To() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (s *Statement) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Statement) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (s *Statement) ReversedLine() graph.Line { _ = "STUB: not implemented"; return *new(graph.Line) }

func ParseNQuad(statement string) (*Statement, error) { _ = "STUB: not implemented"; return nil, nil }

type Decoder struct {
	scanner *bufio.Scanner

	strings store
	ids     map[string]int64
}

func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

func (dec *Decoder) Reset(r io.Reader) { _ = "STUB: not implemented"; return }

func (dec *Decoder) Unmarshal() (*Statement, error) { _ = "STUB: not implemented"; return nil, nil }

func (dec *Decoder) idFor(s string) int64 { _ = "STUB: not implemented"; return 0 }

func (dec *Decoder) Terms() map[string]int64 { _ = "STUB: not implemented"; return nil }

type store map[string]string

func (is store) intern(s string) string { _ = "STUB: not implemented"; return "" }

func escape(lq, s, rq string) string { _ = "STUB: not implemented"; return "" }

func unEscape(r []rune) string { _ = "STUB: not implemented"; return "" }
