//go:build !safe
// +build !safe

package iterator

import (
	"unsafe"

	"gonum.org/v1/gonum/graph"
)

type mapIter struct {
	m     *emptyInterface
	hiter hiter
}

type emptyInterface struct {
	typ, word unsafe.Pointer
}

func newMapIterNodes(m map[int64]graph.Node) *mapIter { _ = "STUB: not implemented"; return nil }

func newMapIterEdges(m map[int64]graph.Edge) *mapIter { _ = "STUB: not implemented"; return nil }

func newMapIterLines(m map[int64]graph.Line) *mapIter { _ = "STUB: not implemented"; return nil }

func newMapIterWeightedLines(m map[int64]graph.WeightedLine) *mapIter {
	_ = "STUB: not implemented"
	return nil
}

func newMapIterByWeightedEdges(m map[int64]graph.WeightedEdge) *mapIter {
	_ = "STUB: not implemented"
	return nil
}

func newMapIterByLines(m map[int64]map[int64]graph.Line) *mapIter {
	_ = "STUB: not implemented"
	return nil
}

func newMapIterByWeightedLines(m map[int64]map[int64]graph.WeightedLine) *mapIter {
	_ = "STUB: not implemented"
	return nil
}

func eface(i interface{}) *emptyInterface { _ = "STUB: not implemented"; return nil }

func (it *mapIter) id() int64 { _ = "STUB: not implemented"; return 0 }

func (it *mapIter) node() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (it *mapIter) line() graph.Line { _ = "STUB: not implemented"; return *new(graph.Line) }

func (it *mapIter) weightedLine() graph.WeightedLine {
	_ = "STUB: not implemented"
	return *new(graph.WeightedLine)
}

func (it *mapIter) next() bool { _ = "STUB: not implemented"; return false }

//go:linkname mapiterinit runtime.mapiterinit
//go:noescape
func mapiterinit(t, m unsafe.Pointer, it *hiter)

//go:linkname mapiterkey reflect.mapiterkey
//go:noescape
func mapiterkey(it *hiter) (key unsafe.Pointer)

//go:linkname mapiterelem reflect.mapiterelem
//go:noescape
func mapiterelem(it *hiter) (elem unsafe.Pointer)

//go:linkname mapiternext reflect.mapiternext
//go:noescape
func mapiternext(it *hiter)
