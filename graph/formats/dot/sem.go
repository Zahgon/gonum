package dot

import (
	"gonum.org/v1/gonum/graph/formats/dot/ast"
)

func check(file *ast.File) error { _ = "STUB: not implemented"; return nil }

func checkGraph(graph *ast.Graph) error { _ = "STUB: not implemented"; return nil }

func checkStmt(graph *ast.Graph, stmt ast.Stmt) error { _ = "STUB: not implemented"; return nil }

func checkNodeStmt(graph *ast.Graph, stmt *ast.NodeStmt) error {
	_ = "STUB: not implemented"
	return nil
}

func checkEdgeStmt(graph *ast.Graph, stmt *ast.EdgeStmt) error {
	_ = "STUB: not implemented"
	return nil
}

func checkEdge(graph *ast.Graph, from ast.Vertex, to *ast.Edge) error {
	_ = "STUB: not implemented"
	return nil
}

func checkAttrStmt(graph *ast.Graph, stmt *ast.AttrStmt) error {
	_ = "STUB: not implemented"
	return nil
}

func checkAttr(graph *ast.Graph, kind ast.Kind, attr *ast.Attr) error {
	_ = "STUB: not implemented"
	return nil
}

func checkSubgraph(graph *ast.Graph, subgraph *ast.Subgraph) error {
	_ = "STUB: not implemented"
	return nil
}

func checkVertex(graph *ast.Graph, vertex ast.Vertex) error { _ = "STUB: not implemented"; return nil }

func checkNode(graph *ast.Graph, node *ast.Node) error { _ = "STUB: not implemented"; return nil }
