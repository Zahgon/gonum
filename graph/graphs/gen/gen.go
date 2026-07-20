package gen

import (
	"gonum.org/v1/gonum/graph"
)

type GraphBuilder interface {
	HasEdgeBetween(xid, yid int64) bool
	graph.Builder
}

func abs(a int) int { _ = "STUB: not implemented"; return 0 }

type NodeIDGraphBuilder interface {
	graph.Builder
	graph.NodeWithIDer
}

type IDer interface {
	Len() int

	ID(int) int64
}

type IDRange struct{ First, Last int64 }

func (r IDRange) Len() int       { _ = "STUB: not implemented"; return 0 }
func (r IDRange) ID(i int) int64 { _ = "STUB: not implemented"; return 0 }

type IDSet []int64

func (s IDSet) Len() int       { _ = "STUB: not implemented"; return 0 }
func (s IDSet) ID(i int) int64 { _ = "STUB: not implemented"; return 0 }

func Complete(dst NodeIDGraphBuilder, ids IDer) { _ = "STUB: not implemented"; return }

func Cycle(dst NodeIDGraphBuilder, cycle IDer) { _ = "STUB: not implemented"; return }

func cycleNoCheck(dst NodeIDGraphBuilder, cycle IDer) { _ = "STUB: not implemented"; return }

func Path(dst NodeIDGraphBuilder, path IDer) { _ = "STUB: not implemented"; return }

func Star(dst NodeIDGraphBuilder, center int64, leaves IDer) { _ = "STUB: not implemented"; return }

func Wheel(dst NodeIDGraphBuilder, center int64, cycle IDer) { _ = "STUB: not implemented"; return }

func Tree(dst NodeIDGraphBuilder, n int, nodes IDer) { _ = "STUB: not implemented"; return }

func check(ids IDer, extra ...int64) error { _ = "STUB: not implemented"; return nil }
