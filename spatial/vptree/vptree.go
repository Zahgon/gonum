package vptree

import (
	"container/heap"
	"errors"
	"math"
	"math/rand/v2"
)

type Comparable interface {
	Distance(Comparable) float64
}

type Point []float64

func (p Point) Distance(c Comparable) float64 { _ = "STUB: not implemented"; return 0 }

type Node struct {
	Point   Comparable
	Radius  float64
	Closer  *Node
	Further *Node
}

type Tree struct {
	Root  *Node
	Count int
}

func New(p []Comparable, effort int, src rand.Source) (t *Tree, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var pointAtInfinity = errors.New("vptree: point at infinity")

type builder struct {
	work []float64
	intn func(n int) int
	shuf func(n int, swap func(i, j int))
}

func (b *builder) build(s []Comparable, effort int) *Node { _ = "STUB: not implemented"; return nil }

func (b *builder) selectVantage(s []Comparable, effort int) Comparable {
	_ = "STUB: not implemented"
	return *new(Comparable)
}

func (b *builder) random(n int, s []Comparable) []Comparable { _ = "STUB: not implemented"; return nil }

func (b *builder) partition(v Comparable, s []Comparable) (radius float64, closer, further []Comparable) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

type byDist struct {
	dists  []float64
	points []Comparable
}

func (c byDist) Len() int           { _ = "STUB: not implemented"; return 0 }
func (c byDist) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (c byDist) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (t *Tree) Len() int { _ = "STUB: not implemented"; return 0 }

var inf = math.Inf(1)

func (t *Tree) Nearest(q Comparable) (Comparable, float64) {
	_ = "STUB: not implemented"
	return *new(Comparable), 0
}

func (n *Node) search(q Comparable, dist float64) (*Node, float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

type ComparableDist struct {
	Comparable Comparable
	Dist       float64
}

type Heap []ComparableDist

func (h *Heap) Max() ComparableDist  { _ = "STUB: not implemented"; return *new(ComparableDist) }
func (h *Heap) Len() int             { _ = "STUB: not implemented"; return 0 }
func (h *Heap) Less(i, j int) bool   { _ = "STUB: not implemented"; return false }
func (h *Heap) Swap(i, j int)        { _ = "STUB: not implemented"; return }
func (h *Heap) Push(x interface{})   { _ = "STUB: not implemented"; return }
func (h *Heap) Pop() (i interface{}) { _ = "STUB: not implemented"; return nil }

type NKeeper struct {
	Heap
}

func NewNKeeper(n int) *NKeeper { _ = "STUB: not implemented"; return nil }

func (k *NKeeper) Keep(c ComparableDist) { _ = "STUB: not implemented"; return }

type DistKeeper struct {
	Heap
}

func NewDistKeeper(d float64) *DistKeeper { _ = "STUB: not implemented"; return nil }

func (k *DistKeeper) Keep(c ComparableDist) { _ = "STUB: not implemented"; return }

type Keeper interface {
	Keep(ComparableDist)
	Max() ComparableDist
	heap.Interface
}

func (t *Tree) NearestSet(k Keeper, q Comparable) { _ = "STUB: not implemented"; return }

func (n *Node) searchSet(q Comparable, k Keeper) { _ = "STUB: not implemented"; return }

type Operation func(Comparable, int) (done bool)

func (t *Tree) Do(fn Operation) bool { _ = "STUB: not implemented"; return false }

func (n *Node) do(fn Operation, depth int) (done bool) { _ = "STUB: not implemented"; return false }
