package dot

import (
	"bytes"
	"regexp"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

type Node interface {
	DOTID() string
}

type Attributers interface {
	DOTAttributers() (graph, node, edge encoding.Attributer)
}

type Porter interface {
	FromPort() (port, compass string)

	ToPort() (port, compass string)
}

type Structurer interface {
	Structure() []Graph
}

type MultiStructurer interface {
	Structure() []Multigraph
}

type Graph interface {
	graph.Graph
	DOTID() string
}

type Multigraph interface {
	graph.Multigraph
	DOTID() string
}

type Subgrapher interface {
	Subgraph() graph.Graph
}

type MultiSubgrapher interface {
	Subgraph() graph.Multigraph
}

func Marshal(g graph.Graph, name, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalMulti(g graph.Multigraph, name, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type printer struct {
	buf bytes.Buffer

	prefix string
	indent string
	depth  int
}

type edge struct {
	inGraph  string
	from, to int64
}

func (p *simpleGraphPrinter) print(g graph.Graph, name string, needsIndent, isSubgraph bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *printer) printFrontMatter(name string, needsIndent, isSubgraph, isDirected, isStrict bool) {
	_ = "STUB: not implemented"
	return
}

func (p *printer) writeNode(n graph.Node) { _ = "STUB: not implemented"; return }

func (p *printer) writePorts(port, cp string) { _ = "STUB: not implemented"; return }

func nodeID(n graph.Node) string { _ = "STUB: not implemented"; return "" }

func graphID(g interface{}, n graph.Node) string { _ = "STUB: not implemented"; return "" }

func (p *printer) writeAttributeList(a encoding.Attributer) { _ = "STUB: not implemented"; return }

var attType = []string{"graph", "node", "edge"}

func (p *printer) writeAttributeComplex(ca Attributers) { _ = "STUB: not implemented"; return }

func (p *printer) newline() { _ = "STUB: not implemented"; return }

func (p *printer) openBlock(b string) { _ = "STUB: not implemented"; return }

func (p *printer) closeBlock(b string) { _ = "STUB: not implemented"; return }

type simpleGraphPrinter struct {
	printer
	visited map[edge]bool
}

type multiGraphPrinter struct {
	printer
	visited map[line]bool
}

type line struct {
	inGraph string
	from    int64
	to      int64
	id      int64
}

func (p *multiGraphPrinter) print(g graph.Multigraph, name string, needsIndent, isSubgraph bool) error {
	_ = "STUB: not implemented"
	return nil
}

func quoteID(s string) string { _ = "STUB: not implemented"; return "" }

func isKeyword(s string) bool { _ = "STUB: not implemented"; return false }

var (
	reIdent   = regexp.MustCompile(`^[a-zA-Z\200-\377_][0-9a-zA-Z\200-\377_]*$`)
	reNumeral = regexp.MustCompile(`^[-]?(\.[0-9]+|[0-9]+(\.[0-9]*)?)$`)
)

func isID(s string) bool { _ = "STUB: not implemented"; return false }

func isHTMLID(s string) bool { _ = "STUB: not implemented"; return false }
