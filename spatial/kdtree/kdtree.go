package kdtree

import (
	"container/heap"
	"math"
)

type Interface interface {
	Index(i int) Comparable

	Len() int

	Pivot(Dim) int

	Slice(start, end int) Interface
}

type Bounder interface {
	Bounds() *Bounding
}

type bounder interface {
	Interface
	Bounder
}

type Dim int

type Comparable interface {
	Compare(Comparable, Dim) float64

	Dims() int

	Distance(Comparable) float64
}

type Extender interface {
	Comparable

	Extend(*Bounding) *Bounding
}

type Bounding struct {
	Min, Max Comparable
}

func (b *Bounding) Contains(c Comparable) bool { _ = "STUB: not implemented"; return false }

type Node struct {
	Point       Comparable
	Plane       Dim
	Left, Right *Node
	*Bounding
}

func (n *Node) String() string { _ = "STUB: not implemented"; return "" }

type Tree struct {
	Root  *Node
	Count int
}

func New(p Interface, bounding bool) *Tree { _ = "STUB: not implemented"; return nil }

func build(p Interface, plane Dim) *Node { _ = "STUB: not implemented"; return nil }

func buildBounded(p bounder, plane Dim, bounding bool) *Node { _ = "STUB: not implemented"; return nil }

func (t *Tree) Insert(c Comparable, bounding bool) { _ = "STUB: not implemented"; return }

func (n *Node) insert(c Comparable, d Dim) *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) insertBounded(c Extender, d Dim, bounding bool) *Node {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tree) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *Tree) Contains(c Comparable) bool { _ = "STUB: not implemented"; return false }

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

type Operation func(Comparable, *Bounding, int) (done bool)

func (t *Tree) Do(fn Operation) bool { _ = "STUB: not implemented"; return false }

func (n *Node) do(fn Operation, depth int) (done bool) { _ = "STUB: not implemented"; return false }

func (t *Tree) DoBounded(b *Bounding, fn Operation) bool { _ = "STUB: not implemented"; return false }

func (n *Node) doBounded(fn Operation, b *Bounding, depth int) (done bool) {
	_ = "STUB: not implemented"
	return false
}
