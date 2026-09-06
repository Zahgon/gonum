package dot

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/formats/dot/ast"
)

type AttributeSetters interface {
	DOTAttributeSetters() (graph, node, edge encoding.AttributeSetter)
}

type DOTIDSetter interface {
	SetDOTID(id string)
}

type PortSetter interface {
	SetFromPort(port, compass string) error

	SetToPort(port, compass string) error
}

func Unmarshal(data []byte, dst encoding.Builder) error { _ = "STUB: not implemented"; return nil }

func UnmarshalMulti(data []byte, dst encoding.MultiBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func copyGraph(dst encoding.Builder, src *ast.Graph) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func copyMultigraph(dst encoding.MultiBuilder, src *ast.Graph) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type generator struct {
	directed bool

	ids map[string]graph.Node

	subNodes []graph.Node

	subStart []int

	graphAttr, nodeAttr, edgeAttr encoding.AttributeSetter
}

func (gen *generator) node(dst graph.NodeAdder, id string) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

type simpleGraph struct{ generator }

func (gen *simpleGraph) addStmt(dst encoding.Builder, stmt ast.Stmt) {
	_ = "STUB: not implemented"
	return
}

type basicEdge interface {
	From() graph.Node
	To() graph.Node
}

func applyPortsToEdge(from ast.Vertex, to *ast.Edge, edge basicEdge) {
	_ = "STUB: not implemented"
	return
}

func (gen *simpleGraph) addEdgeStmt(dst encoding.Builder, stmt *ast.EdgeStmt) {
	_ = "STUB: not implemented"
	return
}

func (gen *simpleGraph) addVertex(dst encoding.Builder, v ast.Vertex) []graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (gen *simpleGraph) addEdge(dst encoding.Builder, to *ast.Edge, attrs []*ast.Attr) []graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (gen *generator) pushSubgraph() { _ = "STUB: not implemented"; return }

func (gen *generator) popSubgraph() []graph.Node { _ = "STUB: not implemented"; return nil }

func unique(ns []graph.Node) []graph.Node { _ = "STUB: not implemented"; return nil }

func (gen *generator) isInSubgraph() bool { _ = "STUB: not implemented"; return false }

func (gen *generator) appendSubgraphNode(n graph.Node) { _ = "STUB: not implemented"; return }

type multiGraph struct{ generator }

func (gen *multiGraph) addStmt(dst encoding.MultiBuilder, stmt ast.Stmt) {
	_ = "STUB: not implemented"
	return
}

func (gen *multiGraph) addEdgeStmt(dst encoding.MultiBuilder, stmt *ast.EdgeStmt) {
	_ = "STUB: not implemented"
	return
}

func (gen *multiGraph) addVertex(dst encoding.MultiBuilder, v ast.Vertex) []graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func (gen *multiGraph) addLine(dst encoding.MultiBuilder, to *ast.Edge, attrs []*ast.Attr) []graph.Node {
	_ = "STUB: not implemented"
	return nil
}

func addEdgeAttrs(edge basicEdge, attrs []*ast.Attr) { _ = "STUB: not implemented"; return }

func unquoteID(s string) string { _ = "STUB: not implemented"; return "" }
