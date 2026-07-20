package path

import (
	"math"

	"gonum.org/v1/gonum/floats/scalar"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/mat"
)

type Shortest struct {
	from graph.Node

	nodes []graph.Node

	indexOf map[int64]int

	dist []float64

	next []int

	hasNegativeCycle bool

	negCosts map[negEdge]float64
}

func newShortestFrom(u graph.Node, nodes []graph.Node) Shortest {
	_ = "STUB: not implemented"
	return *new(Shortest)
}

func (p *Shortest) add(u graph.Node) int { _ = "STUB: not implemented"; return 0 }

func (p Shortest) set(to int, weight float64, mid int) { _ = "STUB: not implemented"; return }

func (p Shortest) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (p Shortest) WeightTo(vid int64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Shortest) To(vid int64) (path []graph.Node, weight float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

type ShortestAlts struct {
	from graph.Node

	nodes []graph.Node

	indexOf map[int64]int

	dist []float64

	next [][]int

	hasNegativeCycle bool

	negCosts map[negEdge]float64
}

func newShortestAltsFrom(u graph.Node, nodes []graph.Node) ShortestAlts {
	_ = "STUB: not implemented"
	return *new(ShortestAlts)
}

func (p *ShortestAlts) add(u graph.Node) int { _ = "STUB: not implemented"; return 0 }

func (p ShortestAlts) set(to int, weight float64, mid int) { _ = "STUB: not implemented"; return }

func (p ShortestAlts) addPath(to, mid int) { _ = "STUB: not implemented"; return }

func (p ShortestAlts) From() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (p ShortestAlts) WeightTo(vid int64) float64 { _ = "STUB: not implemented"; return 0 }

func (p ShortestAlts) To(vid int64) (path []graph.Node, weight float64, unique bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

func (p ShortestAlts) AllTo(vid int64) (paths [][]graph.Node, weight float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (p ShortestAlts) AllToFunc(vid int64, fn func(path []graph.Node)) {
	_ = "STUB: not implemented"
	return
}

func (p ShortestAlts) allTo(from, to int, seen []bool, path []graph.Node, fn func(path []graph.Node)) {
	_ = "STUB: not implemented"
	return
}

type negEdge struct{ from, to int }

type AllShortest struct {
	nodes []graph.Node

	indexOf map[int64]int

	dist *mat.Dense

	next [][]int

	forward bool
}

var (
	defaced = scalar.NaNWith(0xdefaced)

	defacedBits = math.Float64bits(defaced)
)

func newAllShortest(nodes []graph.Node, forward bool) AllShortest {
	_ = "STUB: not implemented"
	return *new(AllShortest)
}

func (p AllShortest) at(from, to int) (mid []int) { _ = "STUB: not implemented"; return nil }

func (p AllShortest) set(from, to int, weight float64, mid ...int) {
	_ = "STUB: not implemented"
	return
}

func (p AllShortest) add(from, to int, mid ...int) { _ = "STUB: not implemented"; return }

func (p AllShortest) Weight(uid, vid int64) float64 { _ = "STUB: not implemented"; return 0 }

func (p AllShortest) Between(uid, vid int64) (path []graph.Node, weight float64, unique bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

func (p AllShortest) AllBetween(uid, vid int64) (paths [][]graph.Node, weight float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (p AllShortest) AllBetweenFunc(uid, vid int64, fn func(path []graph.Node)) {
	_ = "STUB: not implemented"
	return
}

func (p AllShortest) allBetween(from, to int, seen []bool, path []graph.Node, fn func([]graph.Node)) {
	_ = "STUB: not implemented"
	return
}

type node int64

func (n node) ID() int64 { _ = "STUB: not implemented"; return 0 }
